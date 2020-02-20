package process
import(
	"fmt"
	"go_code/chatroom/common/message"
	"go_code/chatroom/server/utils"
	"encoding/json"
	"net"
)
type SmsProcess struct{}

func (this *SmsProcess) SendGroupMes(mes message.Message)  {
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data),&smsMes)
	if err != nil{
		fmt.Println("json.Unmarshal(mes.data) err = ",err)
		return
	}

	data,err := json.Marshal(mes) 
	if err != nil{
		fmt.Println("json.Marshal(mes) err = ",err)
		return
	}

	// 遍历所有的用户列表，除了自己以外，一个个发送消息
	for id, up := range userMgr.onlineUsers {
		if id ==  smsMes.UserId {
			continue
		}
		this.SendOnepMes(data,up.Conn)
	}


}
func (this *SmsProcess) SendOnepMes(data []byte, conn net.Conn)  {
	tf := utils.Transter{
		Conn : conn,
	}
	err := tf.WritePkg(data)
	if err != nil{
		fmt.Println("转发消息失败 err = ",err)
		return
	}
}