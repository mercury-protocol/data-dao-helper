package internal

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	LotusPath    string `mapstructure:"LOTUS_PATH"`
	FileSavePath string `mapstructure:"FILE_SAVE_PATH"`
	Mode         string `mapstructure:"MODE"`
}

var DaoConfig *Config

func LoadConfig() {
	viper.AddConfigPath("../")
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Error("Failed to read env file: ", err)
	}

	err = viper.Unmarshal(&DaoConfig)
	if err != nil {
		log.Error("Failed to decode env file: ", err)
	}
}

func ConfigLogger() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})

	switch DaoConfig.Mode {
	case "develop":
		log.SetLevel(log.DebugLevel)
	case "production":
		log.SetLevel(log.InfoLevel)
	}
}
