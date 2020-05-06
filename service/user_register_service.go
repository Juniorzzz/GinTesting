package service

import (
	"GinWeb/model"
	"GinWeb/serializer"
)

// UserRegisterService 管理用户注册服务
type UserRegisterService struct {
	Nickname        string `form:"nickname" json:"nickname" binding:"required,min=2,max=30"`
	Username        string `form:"username" json:"username" binding:"required,min=5,max=30"`
	Telephone       string `form:"telephone" json:"telephone" binding:"required,min=11,max=11"`
	Password        string `form:"password" json:"password" binding:"required,min=8,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=8,max=40"`
}

// Register 用户注册
func (service *UserRegisterService) Register() (model.User, *serializer.Response) {
	user := model.User{
		Username:  service.Username,
		Nickname:  service.Nickname,
		Status:    model.Active,
		Telephone: service.Telephone,
	}

	// 创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return user, &serializer.Response{
			Status: 40002,
			Msg:    "注册失败",
		}
	}

	return user, nil
}
