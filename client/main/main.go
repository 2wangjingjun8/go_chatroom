package main

import(
	"fmt"
	"go_code/chatroom/client/process"
)
var (
	userId int
	userPwd string
	userName string
)

func main()  {
	var key int
	for {
		fmt.Println("----------欢迎登陆多人聊天系统----------")
		fmt.Println("----------1 登陆")
		fmt.Println("----------2 注册")
		fmt.Println("----------3 退出")
		fmt.Println("----------请选择（1-3）")
		fmt.Scanf("%d\n",&key)
		switch key {
			case 1:
				fmt.Println("----------登陆----------")
				
				fmt.Println("请输入id")
				fmt.Scanf("%d\n",&userId)
				fmt.Println("请输入密码")
				fmt.Scanf("%s\n",&userPwd)
				up := &process.UserProcess{}
				up.Login(userId, userPwd)
				
			case 2:
				fmt.Println("----------注册----------")
				fmt.Println("请输入id")
				fmt.Scanf("%d\n",&userId)
				fmt.Println("请输入密码")
				fmt.Scanf("%s\n",&userPwd)
				fmt.Println("请输入昵称")
				fmt.Scanf("%s\n",&userName)
				up := &process.UserProcess{}
				up.Register(userId, userPwd, userName)
				
			case 3:
				fmt.Println("----------退出----------")
				return
			default:
				fmt.Println("输入有误，请重新输入。。。")
		}
	}


}