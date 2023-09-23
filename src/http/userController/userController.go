package usercontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/guilhermesants/AppSocialApi/src/services/userService"
)

type UserController struct {
	UserService userService.UserService
}

func NewUserController(userService userService.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (uc *UserController) ObterUsuario(c *gin.Context) {

	id := c.Query("id")
	guid, _ := uuid.Parse(id)

	user, err := uc.UserService.Get(guid)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, user)
}
