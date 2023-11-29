package controller

import (
	"backend/pkg/jsonHelper"
	"backend/services/userService/repository"
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

	goth.UseProviders(
		google.New(id, key, "http://localhost:5000/api/oauth/callback?provider=google"),
	)

	gothic.Store = sessions.NewCookieStore([]byte(viper.GetString("SECRET")))

	return &OauthHandler{
		repo: repo,
	}
}

func (h *OauthHandler) OauthCallback(w http.ResponseWriter, request *http.Request) {
	user, err := gothic.CompleteUserAuth(w, request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonHelper.HttpResponse(user, w)
}

func (h *OauthHandler) OauthLogout(w http.ResponseWriter, request *http.Request) {
	gothic.Logout(w, request)
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func (h *OauthHandler) OauthInit(w http.ResponseWriter, request *http.Request) {
	if user, err := gothic.CompleteUserAuth(w, request); err == nil {
		jsonHelper.HttpResponse(user, w)
	} else {
		gothic.BeginAuthHandler(w, request)
	}
}
