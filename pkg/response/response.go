package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RespStruct struct {
	Code    RespCode    `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SuccessResponse(ctx *gin.Context, data interface{}) {
	Response(ctx, http.StatusOK, SUCCESS, data)
	return
}

func FailedResponse(ctx *gin.Context, err error) {
	Response(ctx, http.StatusOK, ErrorServer, err.Error())
}

func Response(ctx *gin.Context, httpCode int, code RespCode, data interface{}) {
	ctx.JSONP(httpCode, RespStruct{
		Code:    code,
		Message: code.String(),
		Data:    data,
	})
	ctx.Abort()
	return
}
