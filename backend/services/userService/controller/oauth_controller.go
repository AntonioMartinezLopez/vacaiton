package controller

import (
	"backend/pkg/jsonHelper"
	"backend/pkg/logger"
	"backend/services/userService/jwtHelper"
	"backend/services/userService/models"
	"backend/services/userService/repository"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
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

	gothic.Store = sessions.NewCookieStore([]byte(viper.GetString("SECRET")))

	goth.UseProviders(
		google.New(id, key, fmt.Sprintf("http://%s/api/oauth/callback?provider=google", viper.GetString("URL"))),
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

	logger.Info(string(jsonHelper.ServeJson(user)))

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

//	@Summary		Logout process using google oauth
//	@Tags			OAuth
//	@Description	This Endpoint is used to login user via google oauth provider - this endpoint triggers the process and redirects to the google authentication service
//	@ID				logoaut-oauth
//	@Accept			json
//	@Produce		html
//	@Router			/oauth [get]
func (h *OauthHandler) OauthLogout(w http.ResponseWriter, request *http.Request) {
	gothic.Logout(w, request)
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusTemporaryRedirect)
}

//	@Summary		Initiates login process using google oauth
//	@Tags			OAuth
//	@Description	This Endpoint is used to login user via google oauth provider - this endpoint triggers the process and redirects to the google authentication service
//	@ID				login-oauth
//	@Param			provider	query	string	false	"oauth provider"	default(google)
//	@Accept			json
//	@Produce		json
//	@Router			/oauth [get]
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
