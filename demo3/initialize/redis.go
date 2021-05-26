package initialize

import (
	"fmt"
	"github.com/go-redis/redis"
	"go-project/demo3/global"
)

func Redis() {
	redisCfg := global.GB_GONFIG.RedisConfig
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Host,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println("redis connect ping failed, err:", err)
	} else {
		fmt.Println("redis connect ping response:", pong)
		global.GB_REDIS = client
	}
}
