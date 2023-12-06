package controller

import (
	"backend/pkg/jsonHelper"
	"backend/services/userService/models"
	"context"
	"errors"
	"fmt"
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

func (r fakeUserRepository) GetUserByMail(_ string) (user *models.User, err error) {
	return &(fakeUsers[0]), nil
}

func (r fakeUserRepository) RegisterUser(_ *models.RegisterUserInput) (userId int, err error) {
	return 2, nil
}

func TestFakeFunction(t *testing.T) {
	fakeUserArray := createFakeUser()

	fmt.Println("TEST")

	for _, user := range fakeUserArray {
		if user.Id != 1 {
			t.Fatalf("Wrong value in user: %d", user.Id)
		}
	}

	if len(fakeUserArray) == 0 {
		t.Fatal("Error Array empty")
	}
}

func TestGetUserInfoHandler(t *testing.T) {

	t.Run("Should be able to return user information for given authenticated user", func(t *testing.T) {
		// create request and reponse objects from a artificial context
		// request with extended context as it is a protected route where previously authentication was processed
		req := httptest.NewRequest(http.MethodGet, "/api/auth/user", nil)
		ctx := context.WithValue(req.Context(), "user-claims", models.Claims{UserId: "1"})

		// response object mocked by recorder
		w := httptest.NewRecorder()

		// Create user handler
		userHandler := UserHandler{repo: fakeUserRepository{}}

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
		ctx := context.WithValue(req.Context(), "user-claims", models.Claims{UserId: "9"})

		// response object mocked by recorder
		w := httptest.NewRecorder()

		// Create user handler
		userHandler := UserHandler{repo: fakeUserRepository{}}

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
		userInput := models.RegisterUserInput{}
		req := httptest.NewRequest(http.MethodPost, "/api/auth/signup", strings.NewReader(string(jsonHelper.ServeJson(userInput))))

		// response object mocked by recorder
		w := httptest.NewRecorder()

		// Create user handler
		userHandler := UserHandler{repo: fakeUserRepository{}}

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
		userHandler := UserHandler{repo: fakeUserRepository{}}

		// Conduct handler call
		userHandler.CreateUser(w, req)

		// analyze result
		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("Expected status code to be %d, got %d", http.StatusNotFound, res.StatusCode)
		}
	})

}
