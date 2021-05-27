package config

type Config struct {
	RedisConfig RedisConfig `ini:"redis"`
	Mongo       Mongo       `ini:"mongo"`
}
