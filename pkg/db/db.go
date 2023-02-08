package db

import (
	"github.com/gin-gonic/gin"
	"github.com/shixinshuiyou/framework/database"
	"gorm.io/gorm"
	"log"
	"mayo_web/pkg/appconf"
)

var BaseDB *gorm.DB

func InitDbMiddleware(ctx *gin.Context) {
	db, err := appconf.DbConf.InitDB()
	if err != nil {
		log.Fatalf("初始化db失败,err: %v", err)
	}
	BaseDB = db.WithContext(ctx.Request.Context())
	database.AddGormCallbacks(BaseDB, appconf.AppConf.Name)
}
