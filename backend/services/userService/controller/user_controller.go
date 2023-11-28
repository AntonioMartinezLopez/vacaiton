package controller

import (
	"backend/pkg/jsonHelper"
	"backend/pkg/logger"
	"backend/services/userService/models"
	"backend/services/userService/repository"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
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
	logger.Info(userInput.Email)
	user, error := h.repo.GetUserByMail(userInput.Email)

	if error != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Check password and generate JWT
	passwordCheckError := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password))

	if passwordCheckError != nil {
		http.Error(w, passwordCheckError.Error(), http.StatusUnauthorized)
	}

	// Create JWT Token and Sign it
	expirationTime := time.Now().Add(time.Minute * 20)

	claims := &models.Claims{
		UserId: strconv.Itoa(user.Id),
		Email:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signingSecret := viper.GetString("SECRET")
	signedToken, err := token.SignedString([]byte(signingSecret))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Add jwt token as cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   signedToken,
		Expires: expirationTime,
	})

	jsonHelper.HttpResponse(&models.AuthResponse{Status: models.LoggedIn}, w)
}

func (h *UserHandler) LogoutUser(w http.ResponseWriter, request *http.Request) {

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
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

	token, err := jwt.ParseWithClaims(jwtString, claims, func(t *jwt.Token) (any, error) {
		return []byte(viper.GetString("SECRET")), nil
	})

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

	jsonHelper.HttpResponse(&models.AuthResponse{Status: models.LoggedIn}, w)
}
