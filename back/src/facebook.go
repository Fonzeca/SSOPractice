package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Fonzeca/SSOPractice/src/model"
	fb "github.com/huandu/facebook/v2"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

var (
	confFacebook = &oauth2.Config{}
)

func InitFacebookConfig() {
	confFacebook = &oauth2.Config{
		ClientID:     viper.GetString("facebook.client_id"),
		ClientSecret: viper.GetString("facebook.client_secret"),
		RedirectURL:  viper.GetString("redirect-callback") + "?type=facebook",
		Scopes:       viper.GetStringSlice("facebook.scopes"),
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://www.facebook.com/v15.0/dialog/oauth",
			TokenURL: "https://graph.facebook.com/v15.0/oauth/access_token",
		},
	}
}

func LoginFacebook(c echo.Context) error {

	// Redirect user to Google's consent page to ask for permission
	// for the scopes specified above.
	url := confFacebook.AuthCodeURL("state")
	fmt.Println("Redirected to consent page: ", url)

	return c.String(http.StatusOK, url)
}

func LoginFacebookCallback(c echo.Context) error {
	authCode := c.FormValue("code")
	// // Handle the exchange code to initiate a transport.
	tok, err := confFacebook.Exchange(oauth2.NoContext, authCode)
	if err != nil {
		log.Fatal(err)
	}
	fb.SetHttpClient(confFacebook.Client(context.Background(), tok))
	res, err := fb.Get("/me", fb.Params{
		"access_token": tok.AccessToken,
		"redirect":     false,
		"fields":       "email,name,picture{url}",
	})
	if err != nil {
		log.Fatal(err)
	}

	facebookView := model.FacebookResponse{}
	res.Decode(&facebookView)

	userView := model.UserView{
		Email:     facebookView.Email,
		Name:      facebookView.Name,
		LoginType: "facebook",
		PhotoURL:  facebookView.Picture.Data.URL,
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
