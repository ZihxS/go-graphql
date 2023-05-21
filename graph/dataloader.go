package graph

import (
	"context"
	"net/http"
	"time"

	"github.com/ZihxS/go-graphql/graph/model"
	"gorm.io/gorm"
)

type ctxKey string

const userLoaderKey ctxKey = "user_loader_key"

func DataloaderMiddleware(db *gorm.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userloader := UserLoader{
			maxBatch: 100,
			wait:     time.Millisecond,
			fetch: func(ids []string) ([]*model.User, []error) {
				var users []*model.User
				err := db.Where("id IN ?", ids).Find(&users).Error
				return users, []error{err}
			},
		}

		ctx := context.WithValue(r.Context(), userLoaderKey, &userloader)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUserLoader(ctx context.Context) *UserLoader {
	return ctx.Value(userLoaderKey).(*UserLoader)
}
