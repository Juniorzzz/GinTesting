package serializer

import "GinWeb/model"

// User 用户序列化器
type User struct {
	ID        uint   `form:"id" json:"id"`
	Username  string `form:"name" json:"username"`
	Nickname  string `form:"name" json:"nickname"`
	Status    string `json:"status"`
	Telephone string `json:"telephone"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"created_at"`
}

// UserResponse 单个用户序列化
type UserResponse struct {
	Response
	Data User `json:"data"`
}

// BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		ID:        user.ID,
		Username:  user.Username,
		Nickname:  user.Nickname,
		Status:    user.Status,
		Telephone: user.Telephone,
		Avatar:    user.Avatar,
		CreatedAt: user.CreatedAt.Unix(),
	}
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user model.User) UserResponse {
	return UserResponse{
		Data: BuildUser(user),
	}
}
