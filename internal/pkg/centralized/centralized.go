package centralized

import (
	"github.com/spf13/viper"
	"rest-api/internal/pkg/asynq"
	"rest-api/internal/pkg/config"
	"rest-api/internal/pkg/mongo"
)

var EnvConfig config.Environment

func Centralized() {
	err := getConfiguration(&EnvConfig)
	if err != nil {
		panic("1231231")
	}
	mongo.ProviderNoSQL(EnvConfig)
	asynq.InitAsynq(EnvConfig.Redis)
	asynq.InitAsynqServer(EnvConfig.Redis)
}

func getConfiguration(environment *config.Environment) error {
	viper.SetConfigName("config.yaml")
	viper.AddConfigPath("environment")
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(environment); err != nil {
		return err
	}
	return nil
}
