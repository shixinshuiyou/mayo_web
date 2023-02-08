package appconf

import (
	"github.com/go-ini/ini"
	"github.com/shixinshuiyou/framework/database"
	"github.com/shixinshuiyou/framework/trace"
	"log"
	"os"
)

var cfg *ini.File

var (
	AppConf   = &App{}
	DbConf    = &database.DatabaseConfig{}
	TraceConf = &trace.Config{}
)

func init() {
	var err error
	var configPath string
	if IsProduct() {
		configPath = "conf/prod/app.ini"
	} else if IsTest() {
		configPath = "conf/test/app.ini"
	} else {
		configPath = "conf/dev/app.ini"
	}

	_, err = os.Stat(configPath)
	if os.IsNotExist(err) {
		log.Fatalf(" configPath %s not find ", configPath)
	}

	cfg, err = ini.Load(configPath)
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
		return
	}

	mapTo("app", AppConf)
	mapTo("db", DbConf)
	mapTo("trace", TraceConf)
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
