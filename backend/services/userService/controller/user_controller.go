package controller

import (
	"backend/pkg/jsonHelper"
	"backend/services/userService/jwtHelper"
	"backend/services/userService/models"
	"backend/services/userService/repository"
	"encoding/json"
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

	if requestedId != "" {

		user, error := h.repo.GetUserById(requestedId)

		if error != nil {
			http.Error(w, error.Error(), http.StatusNotFound)
			return
		}

		jsonHelper.HttpResponse(user, w)

	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, request *http.Request) {
	userInput := new(models.RegisterUserInput)
	err := json.NewDecoder(request.Body).Decode(&userInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.repo.RegisterUser(userInput)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonHelper.HttpResponse(struct {
		UserId int
	}{UserId: id}, w)
}

func (h *UserHandler) LoginUser(w http.ResponseWriter, request *http.Request) {

	// Validate user input
	userInput := new(models.SignUpUserInput)
	err := json.NewDecoder(request.Body).Decode(&userInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get user info
	user, error := h.repo.GetUserByMail(userInput.Email)

	if error != nil {
		http.Error(w, error.Error(), http.StatusNotFound)
		return
	}

	// Check password and generate JWT
	passwordCheckError := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password))

	if passwordCheckError != nil {
		http.Error(w, passwordCheckError.Error(), http.StatusUnauthorized)
		return
	}

	// Generate JWT token
	jwtError := jwtHelper.CreateJwtToken(w, strconv.Itoa(user.Id), user.Email)

	if jwtError != nil {
		http.Error(w, jwtError.Error(), http.StatusInternalServerError)
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

	// obtain token from cookies
	c, err := request.Cookie("token")

	if err != nil {
		if err == http.ErrNoCookie {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jwtString := c.Value
	claims := &models.Claims{}

	token, err := jwtHelper.CheckTokenValid(jwtString, claims)

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !token.Valid {
		http.Error(w, "Invalid token.", http.StatusUnauthorized)
		return
	}

	// Renew token
	renewalError := jwtHelper.CreateJwtToken(w, claims.UserId, claims.Email)

	if renewalError != nil {
		http.Error(w, renewalError.Error(), http.StatusInternalServerError)
	}

	jsonHelper.HttpResponse(&models.AuthResponse{Status: models.LoggedIn}, w)
}
