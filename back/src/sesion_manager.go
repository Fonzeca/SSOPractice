package main

import (
	"time"

	"github.com/Fonzeca/SSOPractice/src/model"
	"github.com/google/uuid"
	"github.com/muesli/cache2go"
	"github.com/spf13/viper"
)

var (
	ttl            = time.Duration(viper.GetInt("ttl-session")) * time.Minute
	CacheTableName = "sessions"
	CookieName     = "session"
)

func InitSessionConfig() {
	ttl = time.Duration(viper.GetInt("ttl-session")) * time.Minute
}

func SaveSession(user model.UserView) (string, error) {
	cache := cache2go.Cache(CacheTableName)

	guid := uuid.New()

	//Borramos la sesion duplicada de la cache
	cache.Foreach(func(key interface{}, item *cache2go.CacheItem) {
		userInCache := item.Data().(model.UserView)

		if userInCache.Email == user.Email && userInCache.LoginType == user.LoginType {
			cache.Delete(key)
		}
	})

	cache.Add(guid.String(), ttl, user)

	return guid.String(), nil
}

func GetSession(guid string) (model.UserView, error) {
	cache := cache2go.Cache(CacheTableName)

	res, err := cache.Value(guid)
	if err != nil {
		return model.UserView{}, err
	}

	return res.Data().(model.UserView), nil
}
