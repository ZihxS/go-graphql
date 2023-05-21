package graph

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/ZihxS/go-graphql/graph/model"
	"gorm.io/gorm"
)

type ctxKey string

const userLoaderKey ctxKey = "user_loader_key"

func DataloaderMiddleware(db *gorm.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userLoader := UserLoader{
			maxBatch: 100,
			wait:     time.Millisecond,
			fetch: func(ids []string) ([]*model.User, []error) {
				var users []*model.User
				err := db.Where("id IN ?", ids).Find(&users).Error

				if err != nil {
					return nil, []error{err}
				}

				u := make(map[string]*model.User, len(users))

				for _, user := range users {
					u[fmt.Sprint(user.ID)] = user
				}

				result := make([]*model.User, len(ids))

				for i, id := range ids {
					result[i] = u[id]
				}

				return result, nil
			},
		}

		ctx := context.WithValue(r.Context(), userLoaderKey, &userLoader)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUserLoader(ctx context.Context) *UserLoader {
	return ctx.Value(userLoaderKey).(*UserLoader)
}
