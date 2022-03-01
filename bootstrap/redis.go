package bootstrap

import (
	"github.com/gomodule/redigo/redis"
	"goAdmin/global"
	"time"
)

func InitRedis() {
	redisPool := &redis.Pool{
		MaxIdle:   global.App.Config.Redis.MaxIdle,
		MaxActive: global.App.Config.Redis.MaxActive,
		Dial: func() (redis.Conn, error) {
			rc, err := redis.Dial("tcp", global.App.Config.Redis.Addr,
				redis.DialPassword(global.App.Config.Redis.Password),
				redis.DialDatabase(global.App.Config.Redis.DefaultDB))
			if err != nil {
				return nil, err
			}
			return rc, nil
		},
		IdleTimeout: time.Second * time.Duration(global.App.Config.Redis.IdleTimeout),
		Wait:        global.App.Config.Redis.Wait,
	}

	global.App.PoolRedis = redisPool
}
