
package dao

import (
	"mayo_web/models"
	"gorm.io/gorm/clause"
)

type MayoUserDao struct {
	db *gorm.DB
}

func NewMayoUserDao() *MayoUserDao {
	mayoUser  := MayoUserDao{}
	return &mayoUser
}


func (mayoUserDao *MayoUserDao)GetByID(id int)(mayoUser *models.MayoUser, err error)  {
	err = mayoUserDao.db.Where(" id = ?",  id).First(&mayoUser).Error
	return
}
   
func (mayoUserDao *MayoUserDao)AddOrEdit(mayoUser *models.MayoUser)(err error)  {
	err = mayoUserDao.db.Omit(clause.Associations).Create(&mayoUser).Error
	return
}

func (mayoUserDao *MayoUserDao)Update(old *models.MayoUser, mayoUser *models.MayoUser)(err error)  {
	err = mayoUserDao.db.Model(&old).Omit(clause.Associations).Save(&mayoUser).Error
	return
}
     
func (mayoUserDao *MayoUserDao)UpdateWithParams(mayoUser *models.MayoUser,params map[string]interface{})(err error)  {
	err = mayoUserDao.db.Model(&mayoUser).Omit(clause.Associations).Updates(params).Error
	return
}


