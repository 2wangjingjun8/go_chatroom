package model

import (
	"github.com/garyburd/redigo/redis"
	"go_code/chatroom/common/message"
	"fmt"
	"encoding/json"
)

var (
	MyUserDao *UserDao
)
type UserDao struct{
	pool *redis.Pool
}

// 工厂模式，创建一个userDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool:pool,
	}
	return
}
func (this *UserDao) GetUserById(conn redis.Conn,id int)  (user *User, err error)  {
	fmt.Println("hah id = ",id)
	res,err := redis.String(conn.Do("HGET","users",id))
	fmt.Printf("hah res = %v ,type = %T",res,res)
	if err != nil {
		// 错误
		if err == redis.ErrNil{ // 表示在users 哈希中，没有找到对应的id
			err = ERROR_USER_NOTEXISTS
		}
		return
	}

	user = &User{}

	//该地方怎么会转化后变为空呢
	err = json.Unmarshal([]byte(res),user)
	fmt.Println("hah user = ",user)

	if err != nil {
		fmt.Println("json.Unmarshal err = ",err)
		return
	}
	return
}

func (this *UserDao) Login(userId int,userPwd string) (user *User, err error) {
	// 先从UserDao的连接池中取出一根链接
	conn :=this.pool.Get()
	defer conn.Close()
	user,err = this.GetUserById(conn, userId)
	if err != nil {
		return
	}

	// 这时证明这个用户是获取到
	fmt.Printf("真实密码user.UserPwd = %v userPwd = %v\n",user.UserPwd,userPwd)
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}
func (this *UserDao) Register(user *message.User) ( err error) {
	// 先从UserDao的连接池中取出一根链接
	conn :=this.pool.Get()
	defer conn.Close()
	_,err = this.GetUserById(conn,user.UserId)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}
	//这是,说明id在redis还没有，则可以完成注册
	data,err := json.Marshal(user) //序列化
	if err != nil {
		return
	}
	//入库
	_,err = conn.Do("hset","users",user.UserId,string(data))
	if err != nil {
		fmt.Println("保存注册用户错误 err = ", err)
	}
	return
}
