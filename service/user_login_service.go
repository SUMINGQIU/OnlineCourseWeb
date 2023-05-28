package service

import (
	"singo/model"
	"singo/serializer"
)

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

// setSession 设置session
// func (service *UserLoginService) setSession(c *gin.Context, user model.User) {
// 	s := sessions.Default(c)
// 	s.Clear()
// 	s.Set("user_id", user.ID)
// 	s.Save()
// }

// Login 用户登录函数
func (service *UserLoginService) Login() (model.User, *serializer.Response) {
	var user model.User
	// select * from users where username = ""
	if err := model.DB.Where("user_name = ?", service.UserName).First(&user).Error; err != nil {
		return user, &serializer.Response{
			Status: 40001,
			Msg:    "账号或者密码错误",
		}
	}

	if user.CheckPassword(service.Password) == false {
		return user, &serializer.Response{
			Status: 40001,
			Msg:    "账号或者密码错误",
		}
	}

	// 设置session
	// service.setSession(service, user)

	return user, nil
}
