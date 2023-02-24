package service

import (
	"github.com/google/uuid"
	"mayo_web/dao"
	"mayo_web/form"
	"mayo_web/models"
)

type MayoUserService struct {
	mayoUserDao *dao.MayoUserDao
}

func NewMayoUserService() *MayoUserService {
	return &MayoUserService{
		mayoUserDao: dao.NewMayoUserDao(),
	}
}

// 展示
func (mayoUserService *MayoUserService) Detail(id int) (mayoUser *models.MayoUser, err error) {
	return mayoUserService.mayoUserDao.GetByID(id)
}

// 删除
func (mayoUserService *MayoUserService) Delete(id int) bool {
	return false
}

// 登录校验
func (mayoUserService *MayoUserService) CheckUserPassword(loginForm form.UserLoginForm) (string, bool) {
	user, _ := mayoUserService.mayoUserDao.GetByUsername(loginForm.Username)
	if user.Password == loginForm.Password {
		token, _ := uuid.NewUUID()
		return token.String(), true
	}
	return "", false
}

// 更新
func (mayoUserService *MayoUserService) Update(id int) {

}
