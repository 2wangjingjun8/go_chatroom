package process

import(
	"fmt"
	"net"
	"go_code/chatroom/client/utils"
	"go_code/chatroom/common/message"
	"os"
	"encoding/json"
)

func ShowMenu()  {
	for{
		fmt.Println("----------欢迎XX登录----------")
		fmt.Println("----------1. 显示在线好友列表----------")
		fmt.Println("----------2. 发送消息----------")
		fmt.Println("----------3. 信息列表----------")
		fmt.Println("----------4. 退出登录----------")
		fmt.Println("----------请选择（1-4）----------")
		var key int
		var content string

		var smsProcess = SmsProcess{}
		fmt.Scanf("%d\n",&key)
		switch key {
			case 1:
				// fmt.Println("----------显示在线好友列表----------")
				outPutOnlinerUser()
			case 2:
				fmt.Println("请输入您想对大家说的话：")
				fmt.Scanf("%s\n",&content)
				smsProcess.SendGroupMes(content)

			case 3:
				fmt.Println("----------信息列表----------")
			case 4:
				fmt.Println("----------退出登录----------")
				return
		}
	}
}

func KeepConn(conn net.Conn)  {
	transter := &utils.Transter{
		Conn:conn,
	}
	for{
		// fmt.Println("客户端正在等待读取服务端发送消息...")
		mes,err := transter.ReadPkg()
		if err != nil {
			fmt.Println("transter.ReadPkg err = ",err)
			os.Exit(0)
		}

		switch mes.Type {
			case message.NotifyUserStatusMesType:
				//取出data，反序列化
				var notifyUserStatusMes message.NotifyUserStatusMes
				json.Unmarshal([]byte(mes.Data),&notifyUserStatusMes)
				updateUserStatus(&notifyUserStatusMes)
			case message.SmsMesType:
				outPutGroupMes(&mes)
			default:
				fmt.Println("服务器返回了未知的消息类型")
		}
		// fmt.Println("mes = ",mes)
	}

}