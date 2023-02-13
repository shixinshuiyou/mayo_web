
package service

import (
	"mayo_web/dao"
)

type MayoUserService struct {
	mayoUserDao    dao.MayoUserDao
}

func NewMayoUserService() *MayoUserService {
	return &MayoUserService{}
}

// 展示
func (mayoUserService *MayoUserService) Detail() {

}

// 删除
func (mayoUserService *MayoUserService) Delete(id int) bool {
	return false
}

// 列表
func (mayoUserService *MayoUserService) List() {

}

// 更新
func (mayoUserService *MayoUserService) Update(id int) {

}


