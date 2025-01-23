package asynq

import (
	"fmt"
	"github.com/arsyadarmawan/asynq-distributed-task/client"
	"github.com/arsyadarmawan/rest-api/internal/pkg/config"
)

var AsynqClient *client.Client

func InitAsynq(redis config.Redis) *client.Client {
	addr := fmt.Sprintf("%s:%d", redis.Host, redis.Port)
	c := client.InitConfiguration(client.RedisConnection{
		Addr: addr,
		DB:   redis.Db,
	})
	AsynqClient = c
	return c
}
