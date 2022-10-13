package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func main() {
	InitViperConfig()
	InitAllConfig()
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     viper.GetStringSlice("allowed-origins"),
		AllowCredentials: true,
	}))
	e.Use(middleware.Logger())
	e.Debug = true

	e.GET("/google", LoginGoogle)
	e.GET("/callback", LoginGoogleCallback)

	e.GET("/facebook", LoginFacebook)
	e.GET("/facebook_callback", LoginFacebookCallback)

	e.GET("/profile", GetProfile)

	e.Logger.Fatal(e.Start(":3000"))
}

func InitViperConfig() {
	viper.SetConfigName("config.json")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
}

func InitAllConfig() {
	InitGoogleConfig()
	InitFacebookConfig()
	InitSessionConfig()
}
