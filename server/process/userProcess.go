package process

import(
	"fmt"
	"encoding/json"
	"go_code/chatroom/common/message"
	"go_code/chatroom/server/utils"
	"go_code/chatroom/server/model"
	"net"
)

type UserProcess struct{
	Conn net.Conn
	// 增加一个字段，表示该链接是属于哪个user的
	UserId int
} 

// 在这里我们编写通知其他在线的用户的方法
// userId 要通知其他所有在线的用户，我userId上线了
func (this *UserProcess) NotifyOrtherOnliners(userId int)  {
	// 遍历onlineUsers，然后一个一个发送NotifyuserStatusRes
	for id,up := range userMgr.onlineUsers{
		if id == userId{
			continue
		}
		// 开始通知
		up.NotifyOrtherOnliner(userId)
	}
}

func (this *UserProcess) NotifyOrtherOnliner(userId int)  {
	// 组装要发送的消息
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType
	var notifyUserStatusMes  message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.UserStatus = message.USERONLINE

	// 将notifyUserStatusMes消息序列化
	data,err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("NotifyOrtherOnliner Marshal err = ",err)
		return
	}

	mes.Data = string(data)

	// 对mes再次序列化
	data,err = json.Marshal(mes)
	if err != nil {
		fmt.Println("mes Marshal err = ",err)
		return
	}
	tf := &utils.Transter{
		Conn:this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyOrtherOnliner err = ",err)
		return
	}
}

// 登录
func (this *UserProcess) ServerProcesLogin( mes message.Message ) (err error)  {
	// 1.直接从mes中取出mes.Data，并直接反序列化为LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data),&loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal err = ",err)
	}


	// 2. 声明loginResMes,并完成赋值
	var loginResMes message.LoginResMes

	user,err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		if err ==model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		}else if( err ==model.ERROR_USER_PWD){
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		}else{
			loginResMes.Code = 505
			fmt.Println("服务器内部错误...")
		}
	}else{
		loginResMes.Code = 200
		// 登录成功，就把用户信息放到切片onlineUsers中
		this.UserId = loginMes.UserId
		userMgr.AddOnlineUser(this)

		//通知其他在线用户，我上线了
		this.NotifyOrtherOnliners(loginMes.UserId)

		// 先把当前id加到在线用户列表，再返回用户列表信息
		for id,_ := range userMgr.onlineUsers{
			loginResMes.UserIds = append(loginResMes.UserIds,id)
		}

		fmt.Println(user,"登录成功")
	}

	// 3.loginResMes 序列化
	data,err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal err = ",err)
	}
	// 4.声明resMes,将data赋值给resMes.Data
	var resMes message.Message
	resMes.Type = "LoginResMesType"
	resMes.Data = string(data)
	// 5.将resData序列化，发送
	data,err = json.Marshal(resMes) 
	if err != nil {
		fmt.Println("json.Marshal err = ",err)
	}

	// 导入utils包，引入Conn
	transter := &utils.Transter{
		Conn:this.Conn,
	}
	transter.WritePkg(data)
	return
}

// 注册
func (this *UserProcess) ServerProcesRegister(mes message.Message) (err error) {
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data),&registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal err = ",err)
	}
	// 声明registerResMes,并完成赋值
	var registerResMes message.RegisterResMes

	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		}else{
			registerResMes.Code = 506
			registerResMes.Error = "发生未知错误"
		}
	}else{
		registerResMes.Code = 200
	}
	// 3.registerResMes 序列化
	data,err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal err = ",err)
	}
	// 4.声明resMes,将data赋值给resMes.Data
	var resMes message.Message
	resMes.Type = "RegisterResMesType"
	resMes.Data = string(data)
	// 5.将resData序列化，发送
	data,err = json.Marshal(resMes) 
	if err != nil {
		fmt.Println("json.Marshal err = ",err)
	}

	// 导入utils包，引入Conn
	transter := &utils.Transter{
		Conn:this.Conn,
	}
	transter.WritePkg(data)

	return
}