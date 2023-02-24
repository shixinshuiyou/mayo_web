package dao

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"mayo_web/models"
	"mayo_web/pkg/db"
)

type MayoUserDao struct {
	db *gorm.DB
}

func NewMayoUserDao() *MayoUserDao {
	mayoUser := MayoUserDao{
		db: db.BaseDB,
	}
	return &mayoUser
}

func (mayoUserDao *MayoUserDao) GetByID(id int) (mayoUser *models.MayoUser, err error) {
	err = mayoUserDao.db.Where(" id = ?", id).First(&mayoUser).Error
	return
}

func (mayoUserDao *MayoUserDao) Add(mayoUser *models.MayoUser) (err error) {
	err = mayoUserDao.db.Omit(clause.Associations).Create(&mayoUser).Error
	return
}

func (mayoUserDao *MayoUserDao) Update(old *models.MayoUser, mayoUser *models.MayoUser) (err error) {
	err = mayoUserDao.db.Model(&old).Omit(clause.Associations).Save(&mayoUser).Error
	return
}

func (mayoUserDao *MayoUserDao) UpdateWithParams(mayoUser *models.MayoUser, params map[string]interface{}) (err error) {
	err = mayoUserDao.db.Model(&mayoUser).Omit(clause.Associations).Updates(params).Error
	return
}

func (mayoUserDao *MayoUserDao) GetByUsername(username string) (mayoUser *models.MayoUser, err error) {
	err = mayoUserDao.db.Where("username = ? and disable = 1", username).First(&mayoUser).Error
	return
}
