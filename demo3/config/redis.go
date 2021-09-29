package config

type RedisConfig struct {
	Host         string
	Password     string
	MaxIdle      int
	MaxActive    int
	IdleTimeout  int
	PoolSize     int
	ReadTimeout  int
	WriteTimeout int
	DialTimeout  int
	DB           int
}
