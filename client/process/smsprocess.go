package process

import(
	"fmt"
	"go_code/chatroom/common/message"
	"encoding/json"
	"go_code/chatroom/client/utils"
)

type SmsProcess struct{}

func (this *SmsProcess) SendGroupMes(content string) (err error) {
	var mes message.Message
	mes.Type = message.SmsMesType

	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus

	//序列化smsMes
	data,err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("json.Marshal(smsMes) err = ",err.Error())
		return
	}
	mes.Data = string(data)

	//序列化mes
	data,err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(mes) err = ",err.Error())
		return
	}
	tf := utils.Transter{
		Conn : CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("tf.WritePkg(data) err = ",err.Error())
		return
	}
	return
}