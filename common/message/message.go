package message

const(
	LoginMesType = "LoginMes" 
	LoginResMesType = "LoginResMes"
	RegisterMesType = "RegisterMes"
	RegisterResMesType = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType = "SmsMes"
)

// 这里我们定义几个用户状态的常量
const(
	USERONLINE = iota
	USEROFFLINE
	USERBUSY
)

type Message struct{
	Type string `json:"type"` // 消息类型
	Data string `json:"data"` //消息内容
}

type LoginMes struct{
	UserId int `json:"userId"`	// 用户id
	UserPwd string `json:"userPwd"` // 用户密码
	UserName string `json:"userName"` // 用户名字
}

type LoginResMes struct{
	Code int `json:"code"` // 错误状态码
	UserIds []int //增加字段，保存用户id切片
	Error string `json:"error"` // 错误消息
}

type RegisterMes struct{
	User User `json:"user"`
}

type RegisterResMes struct{
	Code int `json:"code"` // 错误状态码
	Error string `json:"error"` // 错误消息
}

// 为了配合服务器推送用户状态变化的信息
type NotifyUserStatusMes struct{
	UserId int `json:"userId"` //用户id
	UserStatus int `json:"userStatus"` //用户状态
}

type SmsMes struct{
	User // 匿名结构体，继承
	Content string `json:"content"` //发送消息的内容
}