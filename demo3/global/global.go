package global

import (
	"go-project/demo3/config"

	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	GB_GONFIG config.Config
	GB_REDIS  *redis.Client
	GB_MONGO  *mongo.Client
)
