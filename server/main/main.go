package main

import(
	"fmt"
	"net"
	"time"
	"go_code/chatroom/server/model"
)

func StartProcess(conn net.Conn )  {
	defer conn.Close()
	// 循环接受客户发来的消息
	// 导入utils包，引入Conn
	processor := &Processor{
		Conn:conn,
	}
	err := processor.ProcessMain()
	if err != nil {
		fmt.Println("客户端服务端协程通讯出错...")
		return
	}


}

// 这里编写了一个函数，完成了对UserDao 的初始化任务
func initUserDao()  {
	// 这里的pool 本身就是一个全局的变量
	// 这里需要注意的是初始化的顺序
	// initPool, 再initUserDao
	model.MyUserDao = model.NewUserDao(pool)

}

func main()  {
	// 当服务器启动时，我们就去初始化我们的redis连接池
	initPool("127.0.0.1:6379",16,0,300 * time.Second)
	initUserDao()
	
	// 提示消息
	fmt.Println("服务器在8889端口监听...")
	listen,err := net.Listen("tcp","0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.Listen err = ",err)
		return
	}
	for{
		fmt.Println("正在等待客户来链接服务器...")
		conn,err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept() err = ",err)
		}
		go StartProcess(conn)
	}
}