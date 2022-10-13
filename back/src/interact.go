package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func GetProfile(c echo.Context) error {
	var cookie *http.Cookie

	cookies := c.Cookies()
	for _, cook := range cookies {
		if cook.Name == CookieName {
			cookie = cook
			break
		}
	}

	user, err := GetSession(cookie.Value)
	if err != nil {
		fmt.Println(err)
		return err
	}

	user.Expires = time.Now().Add(ttl)
	cookie.Expires = user.Expires
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, user)
}
