package models

import (
	"gorm.io/gorm"
)

type MayoUser struct {
	Id       int    `gorm:"column:id" json:"id"`
	UserName string `gorm:"column:user_name" json:"userName"`
}

func (this *MayoUser) TableName() string {
	return "mayo_user"
}

func (*MayoUser) AfterCreate(tx *gorm.DB) (err error) {
	return
}

func (*MayoUser) BeforeCreate(tx *gorm.DB) (err error) {

	return
}
func (*MayoUser) BeforeUpdate(tx *gorm.DB) (err error) {
	/*	if(this.ID > 0){
		changes := Changed(tx)
		if (len(changes) > 0) {
		  	this.LastModified =  time.Now()
		}
	}*/
	return
}
