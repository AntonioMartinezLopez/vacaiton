package jwtHelper

import (
	"backend/pkg/middlewares"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

/*
This function generates a JWT token and sets the generated token as cookie to the response context
*/
func CreateJwtToken(w http.ResponseWriter, userId string, email string) error {

	// Create JWT Token and Sign it
	expirationTime := time.Now().Add(time.Minute * 20)

	claims := &middlewares.Claims{
		UserId: userId,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signingSecret := viper.GetString("SECRET")
	signedToken, err := token.SignedString([]byte(signingSecret))

	if err != nil {
		return err
	}

	// Add jwt token as cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   signedToken,
		Expires: expirationTime,
		Path:    "/",
	})

	return nil
}

func CheckTokenValid(jwtString string, claims *middlewares.Claims) (token *jwt.Token, err error) {
	token, err = jwt.ParseWithClaims(jwtString, claims, func(t *jwt.Token) (any, error) {
		return []byte(viper.GetString("SECRET")), nil
	})
	return
}
