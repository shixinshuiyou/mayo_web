package main

import (
	"github.com/shixinshuiyou/framework/trace"
	"github.com/shixinshuiyou/framework/web/server"
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
			ReadTimeout:  time.Duration(appconf.AppConf.ReadTimeout),
			WriteTimeout: time.Duration(appconf.AppConf.WriteTimeout),
		},
		StatusSrvConf: server.WebServerConfig{
			Host: appconf.AppConf.StatusHost,
			Port: appconf.AppConf.StatusPort,
		},
		TraceConf: trace.Config{
			Host:        appconf.TraceConf.Host,
			Port:        appconf.TraceConf.Port,
			ServiceName: appconf.AppConf.Name,
			Ratio:       appconf.TraceConf.Ratio,
		},
	}

	server := server.NewServer(config)
	server.SetMainRouterFunc(router.InitRouter)

	server.Start()

}
