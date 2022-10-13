package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Fonzeca/SSOPractice/src/model"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"

	googleauth "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

var (
	conf = &oauth2.Config{}
)

func InitGoogleConfig() {
	// Your credentials should be obtained from the Google
	// Developer Console (https://console.developers.google.com).
	conf = &oauth2.Config{
		ClientID:     viper.GetString("google.client_id"),
		ClientSecret: viper.GetString("google.client_secret"),
		RedirectURL:  viper.GetString("redirect-callback") + "?type=google",
		Scopes:       viper.GetStringSlice("google.scopes"),
		Endpoint: oauth2.Endpoint{
			AuthURL:   "https://accounts.google.com/o/oauth2/auth",
			TokenURL:  "https://oauth2.googleapis.com/token",
			AuthStyle: oauth2.AuthStyleInParams,
		},
	}
}

func LoginGoogle(c echo.Context) error {

	// Redirect user to Google's consent page to ask for permission
	// for the scopes specified above.
	url := conf.AuthCodeURL("state")
	fmt.Println("Redirected to consent page: ", url)
	return c.String(http.StatusOK, url)
}

func LoginGoogleCallback(c echo.Context) error {
	authCode := c.FormValue("code")
	// // Handle the exchange code to initiate a transport.
	tok, err := conf.Exchange(oauth2.NoContext, authCode)
	if err != nil {
		log.Fatal(err)
	}

	svc, err := googleauth.NewService(context.Background(), option.WithTokenSource(oauth2.StaticTokenSource(tok)))
	if err != nil {
		log.Fatal(err)
	}
	userinfo, err := svc.Userinfo.Get().Do()
	if err != nil {
		log.Fatal(err)
	}

	userView := model.UserView{
		Email:     userinfo.Email,
		Name:      userinfo.Name,
		LoginType: "google",
		PhotoURL:  userinfo.Picture,
	}
	guid, err := SaveSession(userView)
	if err != nil {
		return err
	}

	userView.Expires = time.Now().Add(ttl)

	cookie := new(http.Cookie)
	cookie.Name = CookieName
	cookie.Value = guid
	cookie.Expires = userView.Expires
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, userView)
}
