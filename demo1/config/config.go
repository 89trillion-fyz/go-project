package config

type Config struct {
	Server           Server          `ini:"server"`
	Mysql            Mysql           `ini:"mysql"`
	MainRedis        Redis           `ini:"main-redis"`
	MainRedisCluster RedisCluster    `ini:"main-redisCluster"`
	TeamRedis        Redis           `ini:"team-redis"`
	Mongo            Mongo           `ini:"mongo"`
	Android          CommonAppConfig `ini:"android"`
	Ios              CommonAppConfig `ini:"ios"`
	V1redis          Redis           `ini:"v1-redis"`
}
