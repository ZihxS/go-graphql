package graph

import "github.com/ZihxS/go-graphql/repo/mysql"

type Resolver struct {
	MeetupsRepo mysql.MeetupsRepo
	UsersRepo   mysql.UsersRepo
}
