package message

type User struct{
	UserId int `json:"userId"`
	UserPwd string `json:"userPwd"`
	UserName string `json:"userName"`
	UserStatus int `json:"userStatus"`//用户的状态
	Sex string `json:"sex"`
}