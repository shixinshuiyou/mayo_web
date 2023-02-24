package form

type UserLoginForm struct {
	Username string `form:"username" binding:"required,alphanum"`
	Password string `form:"password" binding:"required,min=6,max=12"`
}
