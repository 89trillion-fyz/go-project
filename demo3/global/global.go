package global

import (
	"github.com/go-redis/redis"
	"go-project/demo3/config"
)

var (
	GB_GONFIG config.Config
	GB_REDIS  *redis.Client
)
