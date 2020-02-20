package main

import (
	"github.com/garyburd/redigo/redis"
	"time"
)
var pool *redis.Pool
func initPool(address string, maxIdle, maxActive int, idleTimeout time.Duration)  {
	pool = &redis.Pool{
		MaxIdle:maxIdle, //最大空闲连接数
		MaxActive:maxActive, //表示和数据库最大的链接数
		IdleTimeout:idleTimeout, //最大空闲时间
		Dial:func() (redis.Conn, error)  {
			return redis.Dial("tcp",address)
		},
	}
}