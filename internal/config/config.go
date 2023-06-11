package config

import (
	"sync"

	"github.com/ilyakaznacheev/cleanenv"

	"navigation/internal/logging"
)

var config *AppConfig
var once sync.Once

type AppConfig struct {
	isDebug bool `yaml:"is_debug"`

	Storage struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Database string `yaml:"database"`
		Password string `yaml:"password"`
	} `yaml:"storage"`

	User struct {
		Login string `yaml:"login"`
		Password string `yaml:"password"`
	}
}

func GetConfig() *AppConfig {
	return config
}

func init() {
	// This function in once run
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("read app configuration")
		config = &AppConfig{}
		err := cleanenv.ReadConfig("config.yaml", config)
		if err != nil {
			help, _ := cleanenv.GetDescription(config, nil)
			logger.Info(help)
			logger.Fatalln(err)
		}
	})
}
