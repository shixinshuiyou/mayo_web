package models

import (
	"gorm.io/gorm"
	"time"
)

type MayoUser struct {
	Id         int       `gorm:"column:id" json:"id"`
	Username   string    `gorm:"column:username" json:"username"`
	Password   string    `gorm:"column:password" json:"password"`
	Isable     int       `gorm:"column:isable" json:"isable"` // 0-失效
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
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
