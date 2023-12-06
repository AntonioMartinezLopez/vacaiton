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

//	@Summary		Get user related data
//	@Tags			Auth
//	@Description	Get user info (for registered users)
//	@ID				get-user-info
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	models.User
//	@Failure		400	{object}	jsonHelper.HTTPError	"We need ID!!"
//	@Failure		404	{object}	jsonHelper.HTTPError	"Can not find ID"
//	@Router			/auth/user [get]
//	@Security		ApiKeyAuth
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

//	@Summary		Create a new user
//	@Tags			Auth
//	@Description	Register and create a new user
//	@ID				create-user
//	@Accept			json
//	@Produce		json
//	@Param			newUserInput	body		models.RegisterUserInput	true	"User Input for creating a new user"
//	@Success		200				{object}	models.RegisterUserOutput
//	@Failure		400				{object}	jsonHelper.HTTPError	"Invalid input"
//	@Failure		500				{object}	jsonHelper.HTTPError	"Invalid input: User Already exists"
//	@Router			/auth/signup [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, request *http.Request) {

	// Validate Input
	userInput := new(models.RegisterUserInput)
	err := json.NewDecoder(request.Body).Decode(userInput)
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

//	@Summary		Sign in of an user
//	@Tags			Auth
//	@Description	This Endpoint is used to sign in a specific user
//	@ID				login-user
//	@Accept			json
//	@Produce		json
//	@Param			loginUserInput	body		models.SignInUserInput	true	"User Input for login"
//	@Success		200				{object}	models.AuthResponse
//	@Failure		400				{object}	jsonHelper.HTTPError	"Invalid input"
//	@Failure		401				{object}	jsonHelper.HTTPError	"Invalid input: Invalid password
//	@Failure		404				{object}	jsonHelper.HTTPError	"Invalid input: User not found"
//	@Failure		500				{object}	jsonHelper.HTTPError	"Invalid input: User Already exists"
//	@Router			/auth/login [post]
func (h *UserHandler) LoginUser(w http.ResponseWriter, request *http.Request) {

	// Validate user input
	userInput := new(models.SignInUserInput)
	err := json.NewDecoder(request.Body).Decode(userInput)
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

//	@Summary		Log out of an user
//	@Tags			Auth
//	@Description	This Endpoint is used to logout a specific user and delete the corresponding session cookie
//	@ID				logut-user
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	models.AuthResponse
//	@Router			/auth/logout [get]
//	@Security		OAuth2Application[write, admin]
func (h *UserHandler) LogoutUser(w http.ResponseWriter, request *http.Request) {

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Path:    "/",
		Expires: time.Now(),
	})

	jsonHelper.HttpResponse(&models.AuthResponse{Status: models.LoggedOut}, w)
}

//	@Summary		Check Token validity
//	@Tags			Auth
//	@Description	This Endpoint is used to check token in cookie header and renews it
//	@ID				check-token
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	models.AuthResponse
//	@Failure		400	{object}	jsonHelper.HTTPError	"Invalid input"
//	@Failure		401	{object}	jsonHelper.HTTPError	"Invalid input: Invalid password
//	@Failure		404	{object}	jsonHelper.HTTPError	"Invalid input: User not found"
//	@Failure		500	{object}	jsonHelper.HTTPError	"Invalid input: User Already exists"
//	@Router			/auth [get]
//	@Security		OAuth2Application[write, admin]
func (h *UserHandler) CheckTokenValid(w http.ResponseWriter, request *http.Request) {

	userClaims := request.Context().Value("user-claims").(models.Claims)

	// Renew token
	renewalError := jwtHelper.CreateJwtToken(w, userClaims.UserId, userClaims.Email)

	if renewalError != nil {
		jsonHelper.HttpErrorResponse(w, http.StatusInternalServerError, renewalError)
	}

	jsonHelper.HttpResponse(&models.AuthResponse{Status: models.LoggedIn}, w)
}
