package asynq

import (
	"github.com/arsyadarmawan/asynq-distributed-task/server"
	"github.com/arsyadarmawan/rest-api/internal/pkg/config"
	"github.com/hibiken/asynq"
)

var AsynqServer *server.Server

func InitAsynqServer(redis config.Redis) *server.Server {
	srv := server.InitServer(server.RedisServerConfig{
		Address: "localhost:6379",
		DB:      0,
	}, asynq.Config{
		Queues: map[string]int{
			"critical": 6,
			"default":  3,
			"low":      1,
		},
	})
	AsynqServer = srv
	return srv
}
