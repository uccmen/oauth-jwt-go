package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"

	"time"
)

type FBUserInfo struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	FBID  string `json:"id,omitempty"`
}

type JWT struct {
	Token  string `json:"token,omitempty"`
	Expiry string `json:"expiry,omitempty"`
}

func loginCallbackHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}

	state := strings.TrimSpace(r.FormValue("state"))
	if state != oauthStateString {
		log.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	code := strings.TrimSpace(r.FormValue("code"))
	token, err := oauthConf.Exchange(ctx, code)
	if err != nil {
		log.Printf("oauthConf.Exchange() failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	resp, err := http.Get("https://graph.facebook.com/me?access_token=" +
		url.QueryEscape(token.AccessToken) + "&fields=email,name")
	if err != nil {
		log.Printf("Get: %s\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	defer resp.Body.Close()

	var fbUserInfo *FBUserInfo
	err = json.NewDecoder(resp.Body).Decode(&fbUserInfo)
	if err != nil {
		log.Printf("json decode: %s\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	customClaims := CustomJWTClaims{
		Sub:   fbUserInfo.FBID,
		Name:  fbUserInfo.Name,
		Email: fbUserInfo.Email,
	}

	jwtToken, err := createJWTTokenWithOAuthToken(*token, customClaims)
	if err != nil {
		log.Printf("createJWTTokenWithOAuthToken: %s\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	cookie := http.Cookie{Name: "oauthtoken", Value: jwtToken, Expires: token.Expiry}
	http.SetCookie(w, &cookie)

	jwt := JWT{
		Token:  jwtToken,
		Expiry: token.Expiry.Format(time.RFC3339Nano),
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(jwt)
	if err != nil {
		log.Printf("encode: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
	}
}
