package main

import (
	"net/http"

	"github.com/guilhermesants/AppSocialApi/src/config"
	"github.com/guilhermesants/AppSocialApi/src/context/database"
	router "github.com/guilhermesants/AppSocialApi/src/http"
	userrepository "github.com/guilhermesants/AppSocialApi/src/repositories/userRepository"
	userservice "github.com/guilhermesants/AppSocialApi/src/services/userService"
)

func main() {
	config, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	database.ConnectDB(config.Database)
	usuarioRepository := userrepository.New(database.DB)
	usuarioService := userservice.NewUserService(usuarioRepository)

	httpRouter := router.NewHTTPHandler(usuarioService)
	errorServer := http.ListenAndServe(":"+config.Port, httpRouter)
	if err != nil {
		panic(errorServer)
	}
}
