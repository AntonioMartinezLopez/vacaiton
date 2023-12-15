package controller_test

import (
	"backend/pkg/jsonHelper"
	"backend/pkg/middlewares"
	"backend/services/userService/controller"
	"backend/services/userService/models"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Mocking the repository layer with the UserRepoInterface

func createFakeUser() []models.User {

	// create password hash
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("test"), bcrypt.DefaultCost)

	newUserArray := new([]models.User)
	test := append(*newUserArray, models.User{
		Id:        1,
		Firstname: "Testuser",
		Lastname:  "Test",
		Email:     "Test@test.de",
		CreatedAt: &time.Time{},
		UpdatedAt: &time.Time{},
		Password:  string(hashedPassword),
	})

	return test

}

var fakeUsers = createFakeUser()

type fakeUserRepository struct {
}

func (r fakeUserRepository) GetUserById(userid string) (user *models.User, err error) {
	if userid == "1" {
		return &(fakeUsers[0]), nil
	}
	return nil, errors.New("Unknown user id")
}

func (r fakeUserRepository) GetUserByMail(email string) (user *models.User, err error) {
	if email == "Test@test.de" {
		return &(fakeUsers[0]), nil
	}
	return nil, errors.New("ERROR: User not found.")
}

func (r fakeUserRepository) RegisterUser(userInput *models.RegisterUserInput) (userId int, err error) {
	if _, err := r.GetUserByMail(userInput.Email); err == nil {
		return 0, errors.New("ERROR: User with given email already exists.")
	}
	return 2, nil
}

func TestGetUserInfoHandler(t *testing.T) {

	t.Run("Should be able to return user information for given authenticated user", func(t *testing.T) {
		// create request and reponse objects from a artificial context
		// request with extended context as it is a protected route where previously authentication was processed
		req := httptest.NewRequest(http.MethodGet, "/api/auth/user", nil)
		ctx := context.WithValue(req.Context(), "user-claims", middlewares.Claims{UserId: "1"})

		// response object mocked by recorder
		w := httptest.NewRecorder()

		// Create user handler
		userHandler := controller.NewUserHandler(fakeUserRepository{})

		// Conduct handler call
		userHandler.GetUserInfo(w, req.WithContext(ctx))

		// analyze result
		res := w.Result()
		defer res.Body.Close()
		user := models.User{}
		err := jsonHelper.DecodeJSON(res.Body, &user)

		if err != nil {
			t.Errorf("Expected error to be nil, but got: %s", err)
		}

		if user.Id != 1 {
			t.Errorf("Expected user id to be 1 got %d", user.Id)
		}
	})

	t.Run("Should return Error with 404 error code when user is not known", func(t *testing.T) {
		// create request and reponse objects from a artificial context
		// request with extended context as it is a protected route where previously authentication was processed
		req := httptest.NewRequest(http.MethodGet, "/api/auth/user", nil)
		ctx := context.WithValue(req.Context(), "user-claims", middlewares.Claims{UserId: "9"})

		// response object mocked by recorder
		w := httptest.NewRecorder()

		// Create user handler
		userHandler := controller.NewUserHandler(fakeUserRepository{})

		// Conduct handler call
		userHandler.GetUserInfo(w, req.WithContext(ctx))

		// analyze result
		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != http.StatusNotFound {
			t.Errorf("Expected status code to be %d, got %d", http.StatusNotFound, res.StatusCode)
		}
	})

}

func TestCreateUserHandler(t *testing.T) {
	// create request and reponse objects from a artificial context

	// Test normal procedure
	t.Run("Should create user with correct input", func(t *testing.T) {
		userInput := models.RegisterUserInput{
			Firstname: "test",
			Lastname:  "test",
			Email:     "Test2@test.de",
			Password:  "test",
		}
		req := httptest.NewRequest(http.MethodPost, "/api/auth/signup", strings.NewReader(string(jsonHelper.ServeJson(userInput))))

		// response object mocked by recorder
		w := httptest.NewRecorder()

		// Create user handler
		userHandler := controller.NewUserHandler(fakeUserRepository{})

		// Conduct handler call
		userHandler.CreateUser(w, req)

		// analyze result
		res := w.Result()
		defer res.Body.Close()
		registerUserOutput := models.RegisterUserOutput{}
		err := jsonHelper.DecodeJSON(res.Body, &registerUserOutput)

		if err != nil {
			t.Errorf("Expected error to be nil, but got: %s", err)
		}

		if registerUserOutput.UserId != 2 {
			t.Errorf("Expected user id to be 1 got %d", registerUserOutput.UserId)
		}
	})

	// Test procedure for corrupted user input
	t.Run("Should return 400 when user input is corrupted or not correct", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/api/auth/signup", strings.NewReader(string(jsonHelper.ServeJson(""))))

		// response object mocked by recorder
		w := httptest.NewRecorder()

		// Create user handler
		userHandler := controller.NewUserHandler(fakeUserRepository{})

		// Conduct handler call
		userHandler.CreateUser(w, req)

		// analyze result
		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("Expected status code to be %d, got %d", http.StatusNotFound, res.StatusCode)
		}
	})

	// Test procedure when email already exists
	t.Run("Should return 500 when email is already given", func(t *testing.T) {
		userInput := models.RegisterUserInput{
			Firstname: "test",
			Lastname:  "test",
			Email:     "Test@test.de",
			Password:  "test",
		}
		req := httptest.NewRequest(http.MethodPost, "/api/auth/signup", strings.NewReader(string(jsonHelper.ServeJson(userInput))))

		// response object mocked by recorder
		w := httptest.NewRecorder()

		// Create user handler
		userHandler := controller.NewUserHandler(fakeUserRepository{})

		// Conduct handler call
		userHandler.CreateUser(w, req)

		// analyze result
		res := w.Result()
		defer res.Body.Close()

		registerUserOutput := models.RegisterUserOutput{}
		err := jsonHelper.DecodeJSON(res.Body, &registerUserOutput)

		if err != nil {
			t.Errorf("Expected error to be nil, but got: %s", err)
		}

		t.Log(registerUserOutput.UserId)

		if res.StatusCode != http.StatusInternalServerError {
			t.Errorf("Expected status code to be %d, got %d", http.StatusInternalServerError, res.StatusCode)
		}
	})
}

func TestLoginUserHandler(t *testing.T) {
	// Test procedure for corrupted user input
	t.Run("Should return 400 when user input is corrupted or not correct", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/api/auth/login", strings.NewReader(string(jsonHelper.ServeJson(""))))

		// response object mocked by recorder
		w := httptest.NewRecorder()

		// Create user handler
		userHandler := controller.NewUserHandler(fakeUserRepository{})

		// Conduct handler call
		userHandler.LoginUser(w, req)

		// analyze result
		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("Expected status code to be %d, got %d", http.StatusBadRequest, res.StatusCode)
		}
	})

	t.Run("Should return 404 when user email not found", func(t *testing.T) {
		userInput := models.SignInUserInput{Email: "unknown@test.de", Password: "test"}
		req := httptest.NewRequest(http.MethodPost, "/api/auth/login", strings.NewReader(string(jsonHelper.ServeJson(userInput))))

		// response object mocked by recorder
		w := httptest.NewRecorder()

		// Create user handler
		userHandler := controller.NewUserHandler(fakeUserRepository{})

		// Conduct handler call
		userHandler.LoginUser(w, req)

		// analyze result
		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != http.StatusNotFound {
			t.Errorf("Expected status code to be %d, got %d", http.StatusNotFound, res.StatusCode)
		}
	})

	t.Run("Should return 401 when user password is wrong", func(t *testing.T) {
		userInput := models.SignInUserInput{Email: "Test@test.de", Password: "wrongpassword"}
		req := httptest.NewRequest(http.MethodPost, "/api/auth/login", strings.NewReader(string(jsonHelper.ServeJson(userInput))))

		// response object mocked by recorder
		w := httptest.NewRecorder()

		// Create user handler
		userHandler := controller.NewUserHandler(fakeUserRepository{})

		// Conduct handler call
		userHandler.LoginUser(w, req)

		// analyze result
		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != http.StatusUnauthorized {
			t.Errorf("Expected status code to be %d, got %d", http.StatusUnauthorized, res.StatusCode)
		}
	})

	t.Run("Should return 200 and created jwt token in cookie header with correct input", func(t *testing.T) {
		userInput := models.SignInUserInput{Email: "Test@test.de", Password: "test"}
		req := httptest.NewRequest(http.MethodPost, "/api/auth/login", strings.NewReader(string(jsonHelper.ServeJson(userInput))))

		// response object mocked by recorder
		w := httptest.NewRecorder()

		// Create user handler
		userHandler := controller.NewUserHandler(fakeUserRepository{})

		// Conduct handler call
		userHandler.LoginUser(w, req)

		// analyze result
		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			t.Errorf("Expected status code to be %d, got %d", http.StatusOK, res.StatusCode)
		}

		cookieHeader := res.Header.Get("Set-Cookie")
		cookieSplit := strings.Split(cookieHeader, "=")
		if cookieHeader == "" || (len(cookieSplit) != 2 && cookieSplit[0] != "token") {
			t.Errorf("Expected Cookie Header to be set, got '%s'", cookieHeader)
		}
	})
}

func TestLogoutUserHandler(t *testing.T) {

	t.Run("Should remove jet token from cookie header", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/auth/logout", nil)
		req.Header.Set("Set-Cookie", "token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEiLCJlbWFpbCI6IlRlc3RAdGVzdC5kZSIsImV4cCI6MTcwMTk0MTUxNX0.3PxGnmuvguZ_IuMn2HCP_PAykkJnvm8qVCCzKGpTimM")

		// response object mocked by recorder
		w := httptest.NewRecorder()

		// Create user handler
		userHandler := controller.NewUserHandler(fakeUserRepository{})

		// Conduct handler call
		userHandler.LogoutUser(w, req)

		// analyze result
		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			t.Errorf("Expected status code to be %d, got %d", http.StatusOK, res.StatusCode)
		}

		cookieHeader := res.Header.Get("Set-Cookie")
		cookieSplit := strings.Split(cookieHeader, ";")
		token := strings.Split(cookieSplit[0], "=")

		if len(token) == 2 && token[1] != "" {
			t.Errorf("Expected Cookie token to be empty, got '%s'", token[1])
		}
	})
}

func TestCheckTokenUserHandler(t *testing.T) {

	t.Run("Should return 200 when used token is valid and renews it", func(t *testing.T) {
		// request with extended context as it is a protected route where previously authentication was processed
		req := httptest.NewRequest(http.MethodGet, "/api/auth", nil)
		ctx := context.WithValue(req.Context(), "user-claims", middlewares.Claims{UserId: "1", Email: "Test@test.de"})

		// response object mocked by recorder
		w := httptest.NewRecorder()

		// Create user handler
		userHandler := controller.NewUserHandler(fakeUserRepository{})

		// Conduct handler call
		userHandler.CheckTokenValid(w, req.WithContext(ctx))

		// analyze result
		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			t.Errorf("Expected status code to be %d, got %d", http.StatusOK, res.StatusCode)
		}

		cookieHeader := res.Header.Get("Set-Cookie")
		cookieSplit := strings.Split(cookieHeader, "=")
		if cookieHeader == "" || (len(cookieSplit) != 2 && cookieSplit[0] != "token") {
			t.Errorf("Expected Cookie Header to be set, got '%s'", cookieHeader)
		}

	})
}
