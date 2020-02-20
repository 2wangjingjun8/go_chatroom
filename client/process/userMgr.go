package process

import(
	"fmt"
	"go_code/chatroom/common/message"
	"go_code/chatroom/client/model"
)

// 客户端要维护的map
var onlinerUsers map[int]*message.User = make(map[int]*message.User, 10)
var CurUser  model.CurUser // 在用户登录成功后，初始化CurUser


// 在客户端显示所有在线的用户
func outPutOnlinerUser()  {
	//遍历onlinerUsers
	for id, _ := range onlinerUsers {
		fmt.Println("用户Id:\t",id)
	}
}

// 编写一个方法，处理返回的NotifyUserStatusMes
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes)  {
	user,ok := onlinerUsers[notifyUserStatusMes.UserId]
	if !ok {
		//原来没有
		user = &message.User{
			UserId:notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.UserStatus
	onlinerUsers[notifyUserStatusMes.UserId] = user

	// 显示在线用户列表
	outPutOnlinerUser()
}