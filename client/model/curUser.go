package model

import(
	"net"
	"go_code/chatroom/common/message"
)
type CurUser struct{
	Conn net.Conn
	message.User
}