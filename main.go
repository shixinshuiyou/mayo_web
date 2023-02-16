package main

import (
	"github.com/shixinshuiyou/framework/log"
	"github.com/shixinshuiyou/framework/trace"
	"github.com/shixinshuiyou/framework/web/server"
	"github.com/sirupsen/logrus"
	"mayo_web/pkg/appconf"
	"mayo_web/router"
	"time"
)

func main() {

	config := server.Config{
		Mode: server.Mode(appconf.AppConf.LogLevel),
		Name: appconf.AppConf.Name,
		MainSrvConf: server.WebServerConfig{
			Host:         appconf.AppConf.MainHost,
			Port:         appconf.AppConf.MainPort,
			ReadTimeout:  120 * time.Second,
			WriteTimeout: 120 * time.Second,
		},
		StatusSrvConf: server.WebServerConfig{
			Host: appconf.AppConf.StatusHost,
			Port: appconf.AppConf.StatusPort,
		},
		LogConf: log.Config{
			ServerTag:   "",
			Level:       logrus.DebugLevel,
			GrayLogConf: nil,
		},
		TraceConf: trace.Config{
			Host:        appconf.TraceConf.Host,
			Port:        appconf.TraceConf.Port,
			ServiceName: appconf.AppConf.Name,
			Ratio:       appconf.TraceConf.Ratio,
		},
	}

	logrus.Infof("%v", config)
	server := server.NewServer(config)
	server.SetMainRouterFunc(router.InitRouter)

	server.Start()

}
