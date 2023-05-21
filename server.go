package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ZihxS/go-graphql/graph"
	"github.com/ZihxS/go-graphql/repo/mysql"
	"github.com/spf13/viper"
	dmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const defaultPort = "1315"

type envConfigs struct {
	Port        string `mapstructure:"APP_PORT"`
	DBStringDSN string `mapstructure:"STRING_DSN"`
	DBUser      string `mapstructure:"DB_USER"`
	DBPass      string `mapstructure:"DB_PASS"`
	DBHost      string `mapstructure:"DB_HOST"`
	DBPort      string `mapstructure:"DB_PORT"`
	DBName      string `mapstructure:"DB_NAME"`
	DBMaxOC     int    `mapstructure:"DB_MAX_OPEN_CONNECTIONS"`
	DBMaxIC     int    `mapstructure:"DB_MAX_IDLE_CONNECTIONS"`
}

var env *envConfigs

func initEnvConfigs() {
	env = loadEnvVariables()
}

func loadEnvVariables() (cfg *envConfigs) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal(err)
	}

	return
}

func main() {
	initEnvConfigs()

	dsn := fmt.Sprintf(env.DBStringDSN, env.DBUser, env.DBPass, env.DBHost, env.DBPort, env.DBName)
	db, err := gorm.Open(dmysql.Open(dsn), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	sqlDB, err := db.DB()

	if err != nil {
		log.Fatal(err.Error())
	}

	defer sqlDB.Close()

	sqlDB.SetMaxOpenConns(env.DBMaxOC)
	sqlDB.SetMaxIdleConns(env.DBMaxIC)
	sqlDB.SetConnMaxLifetime(time.Minute)

	port := env.Port
	if port == "" {
		port = defaultPort
	}

	addr := "localhost:" + port

	graphConfig := graph.Config{Resolvers: &graph.Resolver{
		MeetupsRepo: mysql.MeetupsRepo{DB: db},
		UsersRepo:   mysql.UsersRepo{DB: db},
	}}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graphConfig))

	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", graph.DataloaderMiddleware(db, srv))

	log.Printf("connect to http://localhost:%s for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(addr, nil))
}
