package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const cookieTag = "mayo"

func LoginMiddleware(ctx *gin.Context) {
	mvCookie, err := ctx.Cookie(cookieTag)
	// TODO 登录鉴权（可用redis换成一个随机字符串UUID）
	if err != nil || mvCookie != "MTY3NTY2Mzg0NXxOd3dBTkZwTlNUVlR" {
		ctx.JSONP(http.StatusOK, "")
		return
	}
	ctx.Next()
}
