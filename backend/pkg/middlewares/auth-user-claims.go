package middlewares

import (
	"backend/pkg/jsonHelper"
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func UserClaims(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// extract user claims from header set by api gateway (did the authentication step)
		userClaims := r.Header.Get("X-Auth-User-Claims")
		if userClaims == "" {
			jsonHelper.HttpErrorResponse(w, http.StatusUnauthorized, errors.New("Missing user claims"))
			return
		}

		// decode user claims
		decodedUserClaims := Claims{}
		err := json.Unmarshal([]byte(userClaims), &decodedUserClaims)

		if err != nil {
			jsonHelper.HttpErrorResponse(w, http.StatusInternalServerError, err)
			return
		}

		// extend context with user claim information
		ctx := context.WithValue(r.Context(), "user-claims", decodedUserClaims)

		// pass to next handler with extended context body
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

type Claims struct {
	UserId string `json:"id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}
