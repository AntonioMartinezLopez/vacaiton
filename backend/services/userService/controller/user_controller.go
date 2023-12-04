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

// GetStringByInt example
//
//	@Summary		Add a new pet to the store
//	@Description	get user info
//	@ID				get-user-info
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	models.User
//	@Failure		400		{object}	jsonHelper.HTTPError	"We need ID!!"
//	@Failure		404		{object}	jsonHelper.HTTPError	"Can not find ID"
//	@Router			/auth/user [get]
//  @Security OAuth2Application[write, admin]
func (h *UserHandler) GetUserInfo(w http.ResponseWriter, request *http.Request) {

	userClaims := request.Context().Value("user-claims").(models.Claims)
	requestedId := userClaims.UserId

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

	userClaims := request.Context().Value("user-claims").(models.Claims)

	// Renew token
	renewalError := jwtHelper.CreateJwtToken(w, userClaims.UserId, userClaims.Email)

	if renewalError != nil {
		jsonHelper.HttpErrorResponse(w, http.StatusInternalServerError, renewalError)
	}

	jsonHelper.HttpResponse(&models.AuthResponse{Status: models.LoggedIn}, w)
}
