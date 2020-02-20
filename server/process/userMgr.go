package process

import (
	"fmt"
)

// 因为UserMgr在所有地方都需要使用到
// 设置为全局变量

var (
	userMgr *UserMgr
)

type UserMgr struct{
	onlineUsers map[int]*UserProcess
} 

func init()  {
	userMgr = &UserMgr{
		onlineUsers:make(map[int]*UserProcess, 1024),
	}
}

// 添加或者更新一个在线用户
func (this *UserMgr) AddOnlineUser(up *UserProcess)  {
	this.onlineUsers[up.UserId] = up
}

//根据id,删除一个离线用户
func (this *UserMgr) DelOnlineUser(userId int)  {
	delete(this.onlineUsers,userId)
}

// 返回所有在线的用户
func (this *UserMgr) GetAllOnlineUsers() ( map[int]*UserProcess ) {
	return this.onlineUsers
}

//根据id返回对应的在线用户
func (this *UserMgr) GetOneOnlineUsers(userId int) (up *UserProcess,err error ) {
	up,ok := this.onlineUsers[userId]
	if !ok {
		err = fmt.Errorf("用户%d不存在",userId)
	}
	return 
}

