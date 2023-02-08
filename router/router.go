package router

import (
	"github.com/gin-gonic/gin"
	"mayo_web/pkg/db"
	"mayo_web/pkg/middleware"
	"net/http"
)

func InitRouter(engine *gin.Engine) {
	// 数据库初始化和链路追踪
	engine.Use(db.InitDbMiddleware)

	userGroup := engine.Group("/user", middleware.LoginMiddleware)
	userGroup.GET("/test1", func(context *gin.Context) {
		var name string
		db.BaseDB.Table("advertiser").Raw("select name from advertiser where id = ?", 387528).First(&name)
		context.JSONP(200, map[string]interface{}{
			"code": 0,
			"mes":  "hello " + name,
		})
	})
	userGroup.GET("/test2", func(context *gin.Context) {
		context.JSONP(200, map[string]interface{}{
			"code": 0,
			"mes":  "hello tt",
		})
	})

	engine.NoRoute(NoRouter())
}

func NoRouter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.WriteHeader(http.StatusNotFound)
		ctx.Writer.Write([]byte("路由不正确!")) // 返回基础页面也行
	}
}
