package serializer

import "wedog/model"

// User 用户序列化器
type User struct {
	ID        uint   `form:"id" json:"id"`
	UserName  string `form:"name" json:"user_name"`
	Nickname  string `form:"name" json:"nickname"`
	Status    string `json:"status"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"created_at"`
}

// UserResponse 单个用户序列化
type UserResponse struct {
	Response
	Data User `json:"data"`
}

// BuildUser 序列化用户
func BuildUser(user model.Letter) User {
	return User{

	}
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user model.Letter) UserResponse {
	return UserResponse{
		Data: BuildUser(user),
	}
}
