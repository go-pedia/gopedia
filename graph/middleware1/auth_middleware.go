package middleware1

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/pkg/errors"
	"github.com/sony-nurdianto/go-pedia/graph/model"
	"github.com/sony-nurdianto/go-pedia/graph/postgres"
)

//CurrentUserKey defintion
const CurrentUserKey = "currentUser"

//AuthMiddleware func
func AuthMiddleware(repo postgres.UserRepo) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, err := parseToken(r)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			claim, ok := token.Claims.(jwt.MapClaims)

			if !ok || !token.Valid {
				next.ServeHTTP(w, r)
				return
			}

			user, err := repo.GetUserByID(claim["jti"].(string))
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			//I Dont Know what this mean
			ctx := context.WithValue(r.Context(), CurrentUserKey, user)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}

}

var authHeaderExtractor = &request.PostExtractionFilter{
	Extractor: request.HeaderExtractor{"Authorization"},
	Filter:    stripBearerPrefixFromToken,
}

func stripBearerPrefixFromToken(token string) (string, error) {

	bearer := "BEARER"

	if len(token) > len(bearer) && strings.ToUpper(token[0:len(bearer)]) == bearer {
		return token[len(bearer)+1:], nil
	}

	return token, nil
}

var authExtractor = &request.MultiExtractor{
	authHeaderExtractor,
	request.ArgumentExtractor{"access_token"},
}

func parseToken(r *http.Request) (*jwt.Token, error) {
	jwtToken, err := request.ParseFromRequest(r, authExtractor, func(token *jwt.Token) (interface{}, error) {
		t := []byte(os.Getenv(("JWT_SECRET")))
		return t, nil
	})

	return jwtToken, errors.Wrap(err, "parse Tokken error")
}

func GetCurrentUserFromCTX(ctx context.Context) (*model.User, error) {

	if ctx.Value(CurrentUserKey) == nil {
		return nil, errors.New("no user in context")
	}

	user, ok := ctx.Value(CurrentUserKey).(*model.User)
	if !ok || user.ID == "" {
		return nil, errors.New("no user in context")
	}

	return user, nil
}
