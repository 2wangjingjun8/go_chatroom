package process

import (
	"fmt"
	"net"
	"go_code/chatroom/common/message"
	"go_code/chatroom/client/utils"
	"encoding/json"
	"os"
)

type UserProcess struct{
	// 暂时不需要字段
}

// 登录
func (this *UserProcess) Login(userId int, userPwd string) (err error) {
	// fmt.Printf("用户id:%v 密码：%v",userId,userPwd)
	// return nil

	// 1.链接到服务器
	conn,err := net.Dial("tcp","127.0.0.1:8889")
	if err != nil {
		fmt.Println("net.Dial err = ",err)
		return
	}
	defer conn.Close()

	// 2.准备通过conn发送消息到服务器
	var mes message.Message
	mes.Type = message.LoginMesType

	// 3.创建LoginMes消息结构体
	var LoginMes message.LoginMes
	LoginMes.UserId = userId
	LoginMes.UserPwd = userPwd
	fmt.Println("LoginMes = ",LoginMes)

	// 4.将LoginMes序列化
	data,err := json.Marshal(LoginMes)
	if err != nil {
		fmt.Println("json.Marshal err = ",err)
		return
	}

	// 5.将data赋给mes.Data
	mes.Data = string(data)

	// 6.将mes进行序列化
	data,err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err = ",err)
		return
	}
	// 到这，data就是我们要发送的数据
	transter := utils.Transter{
		Conn:conn,
	}
	transter.WritePkg(data)

	// 这里还需处理服务端返回来的消息
	mes,err = transter.ReadPkg()
	if err != nil {
		fmt.Println("readPkg()",err)
		return
	}
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data),&loginResMes)
	if err != nil {
		fmt.Println("json.Unmarshal",err)
		return
	}
	if loginResMes.Code == 200 {
		//登录成功
		//初始化CurUser
		CurUser.Conn = conn
		CurUser.UserId = userId
		CurUser.UserStatus = message.USERONLINE
		// 获取在线用户列表
		fmt.Println("当前在线用户列表:")
		for _,v := range loginResMes.UserIds{
			if v == userId {
				continue
			}
			fmt.Println("当前用户id:",v)

			//完成对客户端的onlinerUsers初始化
			user := &message.User{
				UserId : v,
				UserStatus : message.USERONLINE,
			}
			onlinerUsers[v] = user
		}
		fmt.Printf("\n\n")

		fmt.Println("登录成功")
		//开启协程保持conn通讯，显示二级菜单
		go KeepConn(conn)

		ShowMenu()


	}else{
		fmt.Println("登录失败,err = ",loginResMes.Error)
	}

	return

}

// 注册
func (this *UserProcess) Register(userId int, userPwd string, userName string) (err error) {
	// 1.链接到服务器
	conn,err := net.Dial("tcp","127.0.0.1:8889")
	if err != nil {
		fmt.Println("net.Dial err = ",err)
		return
	}
	defer conn.Close()

	// 2.准备通过conn发送消息到服务器
	var mes message.Message
	mes.Type = message.RegisterMesType

	// 3.创建LoginMes消息结构体
	var RegisterMes message.RegisterMes
	RegisterMes.User.UserId = userId
	RegisterMes.User.UserPwd = userPwd
	RegisterMes.User.UserName = userName
	fmt.Println("RegisterMes = ",RegisterMes)

	// 4.将RegisterMes序列化
	data,err := json.Marshal(RegisterMes)
	if err != nil {
		fmt.Println("json.Marshal err = ",err)
		return
	}

	// 5.将data赋给mes.Data
	mes.Data = string(data)

	// 6.将mes进行序列化
	data,err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err = ",err)
		return
	}
	// 到这，data就是我们要发送的数据
	transter := utils.Transter{
		Conn:conn,
	}
	transter.WritePkg(data)

	// 这里还需处理服务端返回来的消息
	mes,err = transter.ReadPkg()
	if err != nil {
		fmt.Println("readPkg()",err)
		return
	}
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data),&registerResMes)
	if err != nil {
		fmt.Println("json.Unmarshal",err)
		return
	}
	if registerResMes.Code == 200 {
		fmt.Println("注册成功,请重新登录吧")
		os.Exit(0)
	}else{
		fmt.Println("注册失败,err = ",registerResMes.Error)
		os.Exit(0)
	}
	return
}