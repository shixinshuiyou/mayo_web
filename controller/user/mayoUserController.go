package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/shixinshuiyou/framework/util/gconv"
	"mayo_web/form"
	"mayo_web/pkg/response"
	"mayo_web/service"
)

func LoginController(c *gin.Context) {
	var userForm form.UserLoginForm
	err := c.Bind(&userForm)
	if err != nil {
		response.FailedResponse(c, err)
		return
	}
	token, flag := service.NewMayoUserService().CheckUserPassword(userForm)
	if !flag {
		response.FailedResponse(c, errors.New("登录失败！"))
		return
	}
	response.SuccessResponse(c, token)
	return
}

func DetailController(c *gin.Context) {
	c.Request.ParseForm()
	id := c.Request.Form.Get("id")

	userService := service.NewMayoUserService()
	user, err := userService.Detail(gconv.Int(id))
	if err != nil {
		response.FailedResponse(c, err)
		return
	}
	response.SuccessResponse(c, user)
	return
}
