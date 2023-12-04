package middleware

import (
	"backend/pkg/jsonHelper"
	"backend/services/userService/jwtHelper"
	"backend/services/userService/models"
	"context"
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt"
)

func JwtGuard(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Obtain token from cookies
		c, err := r.Cookie("token")

		if err != nil {
			if err == http.ErrNoCookie {
				jsonHelper.HttpErrorResponse(w, http.StatusUnauthorized, err)
				return
			}
			jsonHelper.HttpErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		// Validate token
		jwtString := c.Value
		claims := &models.Claims{}
		token, err := jwtHelper.CheckTokenValid(jwtString, claims)

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				jsonHelper.HttpErrorResponse(w, http.StatusUnauthorized, err)
				return
			}
			jsonHelper.HttpErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		if !token.Valid {
			jsonHelper.HttpErrorResponse(w, http.StatusUnauthorized, errors.New("Invalid token."))
			return
		}

		// extend context with user claim information
		ctx := context.WithValue(r.Context(), "user-claims", *claims)

		// pass to next handler with extended context body
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
