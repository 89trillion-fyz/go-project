package global

import (
	"go-project/demo3/config"

	"github.com/go-redis/redis"
)

var (
	GB_GONFIG config.Config
	GB_REDIS  *redis.Client
)
