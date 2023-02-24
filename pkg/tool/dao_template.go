package tool

var daoTemplate = `
package dao

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"mayo_web/models"
)

type {{class}}Dao struct {
	db *gorm.DB
}

func New{{class}}Dao() *{{class}}Dao {
	{{obj}}  := {{class}}Dao{}
	return &{{obj}}
}


func ({{obj}}Dao *{{class}}Dao)GetByID(id int)({{obj}} *models.{{class}}, err error)  {
	err = {{obj}}Dao.db.Where(" id = ?",  id).First(&{{obj}}).Error
	return
}
   
func ({{obj}}Dao *{{class}}Dao)Add({{obj}} *models.{{class}})(err error)  {
	err = {{obj}}Dao.db.Omit(clause.Associations).Create(&{{obj}}).Error
	return
}

func ({{obj}}Dao *{{class}}Dao)Update(old *models.{{class}}, {{obj}} *models.{{class}})(err error)  {
	err = {{obj}}Dao.db.Model(&old).Omit(clause.Associations).Save(&{{obj}}).Error
	return
}
     
func ({{obj}}Dao *{{class}}Dao)UpdateWithParams({{obj}} *models.{{class}},params map[string]interface{})(err error)  {
	err = {{obj}}Dao.db.Model(&{{obj}}).Omit(clause.Associations).Updates(params).Error
	return
}


`
