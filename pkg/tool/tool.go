package tool

import (
	"fmt"
	"github.com/alexkappa/mustache"
	"mayo_web/pkg/appconf"
	"os"
	"strings"
)

func GenModel(table string, dbType string) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		appconf.DbConf.User,
		appconf.DbConf.Password,
		appconf.DbConf.Host,
		appconf.DbConf.Name)
	//if dbType == "log" {
	//	dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
	//		appconf.DbLogSetting.User,
	//		appconf.DbLogSetting.Password,
	//		appconf.DbLogSetting.Host,
	//		appconf.DbLogSetting.Name)
	//}

	path := "./models"
	isNotExistMkDir(path)

	t2t := NewTable2Struct()
	t2t.EnableJsonTag(true)
	// 个性化配置
	t2t.Config(&T2tConfig{
		// 如果字段首字母本来就是大写, 就不添加tag, 默认false添加, true不添加
		RmTagIfUcFirsted: false,
		// tag的字段名字是否转换为小写, 如果本身有大写字母的话, 默认false不转
		TagToLower: false,
		// 字段首字母大写的同时, 是否要把其他字母转换为小写,默认false不转换
		UcFirstOnly: false,
		//// 每个struct放入单独的文件,默认false,放入同一个文件(暂未提供)
		//SeperatFile: false,
	})
	// 开始迁移转换
	t2t.
		// 指定某个表,如果不指定,则默认全部表都迁移
		Table(table).
		// 表前缀
		//Prefix("prefix_").
		// 是否添加json tag
		//EnableJsonTag(true).
		// 生成struct的包名(默认为空的话, 则取名为: package model)
		PackageName("models").
		// tag字段的key值,默认是orm
		TagKey("gorm").
		// 是否添加结构体方法获取表名
		RealNameMethod("TableName").
		// 生成的结构体保存路径
		SavePath(path + "/" + table + ".go").
		// 数据库dsn,这里可以使用 t2t.DB() 代替,参数为 *sql.DB 对象
		Dsn(dsn).
		// 执行
		Run()
}

func GenRepository(table string) {
	class := camelCase(table)
	obj := upperCamelCase(table)

	template := mustache.New()

	template.Parse(strings.NewReader(daoTemplate))
	tmp, _ := template.RenderString(map[string]string{
		"class": class,
		"obj":   obj,
	})
	path := "./dao"
	isNotExistMkDir(path)
	filePath := fmt.Sprintf("%s", path+"/"+table+"_dao.go")
	f, err := os.Create(filePath)
	if err != nil {
		return
	}
	defer f.Close()

	f.WriteString(tmp)
}

func GenService(table string) {
	class := camelCase(table)
	obj := upperCamelCase(table)
	template := mustache.New()

	template.Parse(strings.NewReader(serviceTemplate))
	tmp, _ := template.RenderString(map[string]string{
		"class": class,
		"obj":   obj,
	})
	path := "./service"
	isNotExistMkDir(path)
	filePath := fmt.Sprintf("%s", path+"/"+table+"_service.go")
	f, err := os.Create(filePath)
	if err != nil {
		return
	}
	defer f.Close()

	f.WriteString(tmp)
}

func camelCase(str string) string {
	var text string
	for _, p := range strings.Split(str, "_") {
		// 字段首字母大写的同时, 是否要把其他字母转换为小写
		switch len(p) {
		case 0:
		case 1:
			text += strings.ToUpper(p[0:1])
		default:
			text += strings.ToUpper(p[0:1]) + p[1:]
		}
	}
	return text
}

func upperCamelCase(str string) string {
	var text string
	for _, p := range strings.Split(str, "_") {
		// 字段首字母大写的同时, 是否要把其他字母转换为小写
		switch len(p) {
		case 0:
		case 1:
			text += strings.ToUpper(p[0:1])
		default:
			text += strings.ToUpper(p[0:1]) + p[1:]
		}
	}
	text = strings.ToLower(text[0:1]) + text[1:]
	return text
}

func isNotExistMkDir(path string) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}
