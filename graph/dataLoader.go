package graph

import (
	"context"
	"net/http"

	"time"

	"github.com/go-pg/pg/v9"
	"github.com/sony-nurdianto/go-pedia/graph/model"
)

const userLoaderKey = "userloader"

//DataLoaderMiddlerware Handler
func DataLoaderMiddlerware(db *pg.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userloader := UserLoader{
			maxBatch: 100,
			wait:     1 * time.Millisecond,
			fetch: func(ids []string) ([]*model.User, []error) {
				var users []*model.User

				err := db.Model(&users).Where("id in (?)", pg.In(ids)).Select()
				if err != nil {
					return nil, []error{err}
				}
				u := make(map[string]*model.User, len(users))

				for _, user := range users {
					u[user.ID] = user
				}

				result := make([]*model.User, len(ids))

				for i, id := range ids {
					result[i] = u[id]
				}

				return users, nil
			},
		}

		ctx := context.WithValue(r.Context(), userLoaderKey, &userloader)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUserLoader(ctx context.Context) *UserLoader {
	return ctx.Value(userLoaderKey).(*UserLoader)
}
