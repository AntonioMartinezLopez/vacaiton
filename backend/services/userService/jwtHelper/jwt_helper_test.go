package jwtHelper_test

import (
	"backend/pkg/jsonHelper"
	"backend/services/userService/jwtHelper"
	"backend/services/userService/middleware"
	"backend/services/userService/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestJwtGuard(t *testing.T) {

	t.Run("Should return 401 when no cookie was set", func(t *testing.T) {

		req := &http.Request{}
		w := httptest.NewRecorder()

		// call middleware with missing jwt in cookies
		jwtGuard := middleware.JwtGuard(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userClaims := r.Context().Value("user-claims").(models.Claims)
			jwtHelper.CreateJwtToken(w, userClaims.UserId, userClaims.Email)
		}))
		jwtGuard.ServeHTTP(w, req)

		// analyze result
		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != http.StatusUnauthorized {
			errorMessage := jsonHelper.HTTPError{}
			jsonHelper.DecodeJSON(res.Body, &errorMessage)
			t.Errorf("Expected status code to be %d, got %d", http.StatusUnauthorized, res.StatusCode)
			t.Errorf("Error %s", errorMessage.Message)
		}

	})

	t.Run("Should return 401 when invalid cookie with wrong signature was set", func(t *testing.T) {

		// generate valid token
		tokenWriter := httptest.NewRecorder()
		jwtHelper.CreateJwtToken(tokenWriter, "1", "Test@test.de")
		cookies := strings.Split(tokenWriter.Result().Cookies()[0].Raw, ";")

		// replace something in signature part of JWT
		split := strings.Split(cookies[0], ".")
		signature := []byte(split[2])
		signature[10] = byte('k')
		split[2] = string(signature)
		cookies[0] = strings.Join(split, ".")

		// set invalid cookie
		req := &http.Request{Header: http.Header{"Cookie": cookies}}

		w := httptest.NewRecorder()

		// call middleware with missing jwt in cookies
		jwtGuard := middleware.JwtGuard(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		jwtGuard.ServeHTTP(w, req)

		// analyze result
		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != http.StatusUnauthorized {
			errorMessage := jsonHelper.HTTPError{}
			jsonHelper.DecodeJSON(res.Body, &errorMessage)
			t.Errorf("Expected status code to be %d, got %d", http.StatusUnauthorized, res.StatusCode)
			t.Errorf("Error %s", errorMessage.Message)
		}

	})
	t.Run("Should return 200 when used token is valid and renews it", func(t *testing.T) {

		t.Setenv("SCECRET", "123456")

		// generate valid token
		tokenWriter := httptest.NewRecorder()
		jwtHelper.CreateJwtToken(tokenWriter, "1", "Test@test.de")
		cookie := tokenWriter.Result().Cookies()[0].Raw
		req := &http.Request{Header: http.Header{"Cookie": strings.Split(cookie, ";")}}

		// call Middleware with mocked handler that generates token
		w := httptest.NewRecorder()

		jwtGuard := middleware.JwtGuard(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userClaims := r.Context().Value("user-claims").(models.Claims)
			jwtHelper.CreateJwtToken(w, userClaims.UserId, userClaims.Email)
		}))
		jwtGuard.ServeHTTP(w, req)

		// analyze result
		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			errorMessage := jsonHelper.HTTPError{}
			jsonHelper.DecodeJSON(res.Body, &errorMessage)
			t.Errorf("Expected status code to be %d, got %d", http.StatusOK, res.StatusCode)
			t.Errorf("Error %s", errorMessage.Message)
		}

		cookieHeader := res.Header.Get("Set-Cookie")
		cookieSplit := strings.Split(cookieHeader, "=")
		if cookieHeader == "" || (len(cookieSplit) != 2 && cookieSplit[0] != "token") {
			t.Errorf("Expected Cookie Header to be set, got '%s'", cookieHeader)
		}
	})
}
