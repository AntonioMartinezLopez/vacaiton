package controller

import (
	"backend/pkg/jsonHelper"
	"backend/services/userService/jwtHelper"
	"backend/services/userService/models"
	"backend/services/userService/repository"
	"net/http"

	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"github.com/spf13/viper"
)

type OauthHandler struct {
	repo repository.UserRepo
}

func NewOauthHandler(repo repository.UserRepo) *OauthHandler {

	id := viper.GetString("OAUTH_GOOGLE_ID")
	key := viper.GetString("OAUTH_GOOGLE_KEY")

	goth.UseProviders(
		google.New(id, key, "http://localhost:5000/api/oauth/callback?provider=google"),
	)

	return &OauthHandler{
		repo: repo,
	}
}

func (h *OauthHandler) OauthCallback(w http.ResponseWriter, request *http.Request) {
	user, err := gothic.CompleteUserAuth(w, request)
	if err != nil {
		jsonHelper.HttpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	// Generate JWT token
	jwtError := jwtHelper.CreateJwtToken(w, user.UserID, user.Email)

	if jwtError != nil {
		jsonHelper.HttpErrorResponse(w, http.StatusInternalServerError, jwtError)
		return
	}

	// TO DO: change line and redirect to root page of the application
	// URL should be passed via env variables
	jsonHelper.HttpResponse(&models.AuthResponse{Status: models.LoggedIn}, w)
}

func (h *OauthHandler) OauthLogout(w http.ResponseWriter, request *http.Request) {
	gothic.Logout(w, request)
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func (h *OauthHandler) OauthInit(w http.ResponseWriter, request *http.Request) {
	if user, err := gothic.CompleteUserAuth(w, request); err == nil {

		// Generate JWT token
		jwtError := jwtHelper.CreateJwtToken(w, user.UserID, user.Email)

		if jwtError != nil {
			jsonHelper.HttpErrorResponse(w, http.StatusInternalServerError, jwtError)
			return
		}
		jsonHelper.HttpResponse(&models.AuthResponse{Status: models.LoggedIn}, w)

	} else {
		gothic.BeginAuthHandler(w, request)
	}
}
