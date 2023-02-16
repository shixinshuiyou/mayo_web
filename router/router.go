package router

import (
	//"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"mayo_web/pkg/db"
	"mayo_web/pkg/middleware"
	"mayo_web/pkg/response"
	"net/http"
)

func InitRouter(engine *gin.Engine) {

	//pprof.Register(engine)
	// 数据库初始化和链路追踪
	engine.Use(db.InitDbMiddleware)

	userGroup := engine.Group("/user", middleware.LoginMiddleware)
	userGroup.GET("/test1", func(c *gin.Context) {
		var name string
		db.BaseDB.Table("advertiser").Raw("select user_name from mayo_user where id = ?", 1).First(&name)
		response.SuccessResponse(c, name)
	})
	userGroup.GET("/test2", func(c *gin.Context) {
		response.SuccessResponse(c, map[string]interface{}{
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
