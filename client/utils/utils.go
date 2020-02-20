package utils
import(
	"fmt"
	"net"
	"go_code/chatroom/common/message"
	"encoding/binary"
	"encoding/json"
)

type Transter struct{
	Conn net.Conn
	Buf [8096]byte 
} 

func (this *Transter) ReadPkg() (mes message.Message,err error)  {
	// fmt.Println("读取客户端发送的数据...")
	
	_,err = this.Conn.Read(this.Buf[0:4])
	if err != nil {
		return
	}
	// fmt.Println("读取到数据：",this.Buf[0:4])

	// 切片长度还原uint32
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[0:4])

	//接受发送过来的消息
	n,err := this.Conn.Read(this.Buf[0:pkgLen])
	if n != int(pkgLen) || err != nil {
		return
	}

	err = json.Unmarshal(this.Buf[0:pkgLen],&mes)
	if err != nil {
		fmt.Println("json.Unmarshal err = ",err)
		return
	}
	return
}

func (this *Transter)  WritePkg(data []byte) ( err error) {
	// 先把data的长度发送给服务器
	// 获取data的长度
	var pkgLen uint32
	pkgLen = uint32(len(data))
	// var byt [4]byte
	binary.BigEndian.PutUint32(this.Buf[0:4],pkgLen)
	// 发送长度
	n,err := this.Conn.Write(this.Buf[0:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write err = ",err)
		return
	}
	
	// 发送数据
	n,err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("this.Conn.Write(data) err = ",err)
		return
	}
	return
}