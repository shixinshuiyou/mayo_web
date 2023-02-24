package main

import (
	"flag"
	"mayo_web/pkg/tool"
)

func main() {
	var table string
	var dbType string
	var srv bool
	flag.StringVar(&table, "table", "", "输入table")
	flag.StringVar(&dbType, "dbType", "db", "输入数据库名称dbType")
	flag.BoolVar(&srv, "srv", true, "是否加载service层")
	flag.Parse()
	if table == "" {
		println("请输入表名")
		return
	}
	tool.GenModel(table, dbType)
	tool.GenDao(table)
	tool.GenService(table)

}
