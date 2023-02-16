package controller

import (
	"github.com/gin-gonic/gin"
	"mayo_web/pkg/response"
)

type UserController struct {
}

func (*UserController) LoginController(c *gin.Context) {

	response.SuccessResponse(c, nil)
	return
}
