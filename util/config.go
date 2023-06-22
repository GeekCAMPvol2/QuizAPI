package util

import (
	"github.com/spf13/viper"
)

type Config struct {
	MongoDBUri    string `mapstructure:"MONGO_DB_URI"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	MongoTestUri  string `mapstructure:"MONOG_TEST_URI"`
}

// 環境変数、app.envからデータを読みだす
func Loadenv(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
