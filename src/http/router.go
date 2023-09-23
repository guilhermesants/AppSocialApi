package http

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/guilhermesants/AppSocialApi/src/entities"
	usercontroller "github.com/guilhermesants/AppSocialApi/src/http/userController"
)

func NewHTTPHandler(usuarioService entities.UserRepository) http.Handler {

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true

	router.Use(cors.New(config))
	userController := usercontroller.NewUserController(usuarioService)

	public := router.Group("/service/v1")
	userRegisterGroup := public.Group("/user")
	userRegisterGroup.GET("login", userController.ObterUsuario)

	return router
}
