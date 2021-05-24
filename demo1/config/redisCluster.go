package config

type RedisCluster struct {
	Slot1          string
	Slot1Start     int
	Slot1End       int
	Slot1Slave     string
	Slot2          string
	Slot2Start     int
	Slot2End       int
	Slot2Slave     string
	Slot3          string
	Slot3Start     int
	Slot3End       int
	Slot3Slave     string
	MaxRedirects   int
	ReadOnly       bool
	RouteByLatency bool
	RouteRandomly  bool
	IdleTimeout    int
	ReadTimeout    int
	WriteTimeout   int
	DialTimeout    int
}
