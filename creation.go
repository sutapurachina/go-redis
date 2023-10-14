package repository

import (
	"context"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

var ctx = context.Background()

func HandleError(err error) {
	if err != nil {
		log.Println(err) //todo
	}
}

type RedisRepository struct {
	pool *redis.Pool
}

func NewTradeRedisRepository(config *RedisConfig) *RedisRepository {
	addr := fmt.Sprintf("%s:%s", config.Host, config.Port)
	pool := &redis.Pool{
		MaxIdle:     config.MaxIdle,
		IdleTimeout: time.Duration(config.IdleTimeOut * int(time.Second)),
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", addr)
			return conn, err
		},
	}
	return &RedisRepository{
		pool: pool,
	}
}

func (r *RedisRepository) Close() error {
	err := r.pool.Close()
	return err
}
