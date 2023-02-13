package tool

var serviceTemplate = `
package service

import (
	"mayo_web/dao"
)

type {{class}}Service struct {
	{{obj}}Dao    dao.{{class}}Dao
}

func New{{class}}Service() *{{class}}Service {
	return &{{class}}Service{}
}

// 展示
func ({{obj}}Service *{{class}}Service) Detail() {

}

// 删除
func ({{obj}}Service *{{class}}Service) Delete(id int) bool {
	return false
}

// 列表
func ({{obj}}Service *{{class}}Service) List() {

}

// 更新
func ({{obj}}Service *{{class}}Service) Update(id int) {

}


`
