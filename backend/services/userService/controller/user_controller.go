package controller

import (
	"backend/pkg/jsonHelper"
	"backend/services/userService/jwtHelper"
	"backend/services/userService/models"
	"backend/services/userService/repository"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	repo repository.UserRepo
}

func NewUserHandler(repo repository.UserRepo) *UserHandler {
	return &UserHandler{
		repo: repo,
	}
}

func (h *UserHandler) GetUserInfo(w http.ResponseWriter, request *http.Request) {

	requestedId := chi.URLParam(request, "user_id")

	// Try to retrieve user info based on given id
	if requestedId != "" {
		user, error := h.repo.GetUserById(requestedId)
		if error != nil {
			jsonHelper.HttpErrorResponse(w, http.StatusNotFound, error)
			return
		}
		jsonHelper.HttpResponse(user, w)
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, request *http.Request) {

	// Validate Input
	userInput := new(models.RegisterUserInput)
	err := json.NewDecoder(request.Body).Decode(&userInput)
	if err != nil {
		jsonHelper.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	// Create User
	id, err := h.repo.RegisterUser(userInput)
	if err != nil {
		jsonHelper.HttpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	// Return result
	jsonHelper.HttpResponse(&models.RegisterUserOutput{UserId: id}, w)
}

func (h *UserHandler) LoginUser(w http.ResponseWriter, request *http.Request) {

	// Validate user input
	userInput := new(models.SignUpUserInput)
	err := json.NewDecoder(request.Body).Decode(&userInput)
	if err != nil {
		jsonHelper.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	// Get user info
	user, error := h.repo.GetUserByMail(userInput.Email)

	if error != nil {
		jsonHelper.HttpErrorResponse(w, http.StatusNotFound, error)
		return
	}

	// Check password and generate JWT
	passwordCheckError := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password))

	if passwordCheckError != nil {
		jsonHelper.HttpErrorResponse(w, http.StatusUnauthorized, passwordCheckError)
		return
	}

	// Generate JWT token
	jwtError := jwtHelper.CreateJwtToken(w, strconv.Itoa(user.Id), user.Email)

	if jwtError != nil {
		jsonHelper.HttpErrorResponse(w, http.StatusInternalServerError, jwtError)
		return
	}

	jsonHelper.HttpResponse(&models.AuthResponse{Status: models.LoggedIn}, w)
}

func (h *UserHandler) LogoutUser(w http.ResponseWriter, request *http.Request) {

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Path:    "/",
		Expires: time.Now(),
	})

	jsonHelper.HttpResponse(&models.AuthResponse{Status: models.LoggedOut}, w)
}

func (h *UserHandler) CheckTokenValid(w http.ResponseWriter, request *http.Request) {

	// Obtain token from cookies
	c, err := request.Cookie("token")

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

	// Renew token
	renewalError := jwtHelper.CreateJwtToken(w, claims.UserId, claims.Email)

	if renewalError != nil {
		jsonHelper.HttpErrorResponse(w, http.StatusInternalServerError, renewalError)
	}

	jsonHelper.HttpResponse(&models.AuthResponse{Status: models.LoggedIn}, w)
}
