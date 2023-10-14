package repository

type RedisConfig struct {
	Host        string
	Port        string
	MaxIdle     int
	IdleTimeOut int
}

func NewRedisConfig(host, port string, maxIdle, IdleTimeOut int) *RedisConfig {
	return &RedisConfig{
		Host:        host,
		Port:        port,
		MaxIdle:     maxIdle,
		IdleTimeOut: IdleTimeOut,
	}
}
