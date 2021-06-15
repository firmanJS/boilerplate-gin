package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(".env")
	viper.AddConfigPath("../")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Connont find config file, %s", err)
	}
}

type Config struct {
	GO_PORT       int
	GO_ENV        string
	DB_URI        string
	GO_DEBUG      bool
}

func NewConfig() (defConfig *Config, err error) {
	defConfig = &Config{}

	appPort := viper.GetInt("GO_PORT")
	appEnv := viper.GetString("GO_ENV")
	appDebug := viper.GetBool("GO_DEBUG")
	dbUri := viper.GetString("DB_URI")

	defConfig.GO_ENV = appEnv
	defConfig.GO_PORT = appPort
	defConfig.GO_DEBUG = appDebug
	defConfig.DB_URI = dbUri

	return defConfig, nil
}
