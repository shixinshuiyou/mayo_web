package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"mayo_web/pkg/response"
)

const cookieTag = "mayo"

func LoginMiddleware(ctx *gin.Context) {
	mvCookie, err := ctx.Cookie(cookieTag)
	if err != nil || mvCookie != "MTY3NTY2Mzg0NXxOd3dBTkZwTlNUVlR" {
		response.FailedResponse(ctx, errors.New("请先登录"))
		return
	}
	ctx.Next()
}
