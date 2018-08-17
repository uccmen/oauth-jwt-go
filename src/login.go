package main

import (
	"context"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

var (
	ctx       = context.Background()
	oauthConf = &oauth2.Config{
		ClientID:     os.Getenv("FB_CLIENT_ID"),
		ClientSecret: os.Getenv("FB_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("APP_CALLBACK_REDIRECT_URL"),
		Scopes:       []string{"public_profile"},
		Endpoint:     facebook.Endpoint,
	}
	oauthStateString = uuid()
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}
	url := oauthConf.AuthCodeURL(oauthStateString, oauth2.AccessTypeOnline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
