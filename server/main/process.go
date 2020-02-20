package main

import(
	"net"
	"go_code/chatroom/common/message"
	"go_code/chatroom/server/process"
	"go_code/chatroom/server/utils"
	"io"
	"fmt"
)

type Processor struct{
	Conn net.Conn
}

// 根据客户端发送消息不同，决定调用哪个函数来处理
func (this *Processor) ServerProcesMes(mes message.Message ) (err error) {
	fmt.Println("mes = ",mes)

	switch mes.Type {
		case message.LoginMesType:
			// 处理登录数据消息
			up := &process.UserProcess{
				Conn:this.Conn,
			}
			up.ServerProcesLogin(mes)
		case message.RegisterMesType:
			// 处理注册数据消息
			up := &process.UserProcess{
				Conn:this.Conn,
			}
			up.ServerProcesRegister(mes)
		case message.SmsMesType:
			// 发接受到群聊消息，转发
			smsProcess := process.SmsProcess{}
			smsProcess.SendGroupMes(mes)
			
		default:
			fmt.Println("消息类型不存在，无法法处理数据...")
	}
	return
}

func (this *Processor) ProcessMain() (err error)  {
	for{
		transter := utils.Transter{
			Conn:this.Conn,
		}
		mes,err := transter.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务端也退出。。。")
				
			}else{
				fmt.Println("err = ",err)
			}
			return err
		}
		fmt.Println("mes = ",mes)
		this.ServerProcesMes(mes)

		
	}
}