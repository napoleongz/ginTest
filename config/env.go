package config

import (
	"github.com/spf13/viper"
	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
	"github.com/vrischmann/envconfig"
)


var(
	GlobalConfig Config
)
type Config struct {
	DBConfig struct{
		DBAddr string
		DBPort string
		DBUsername string
		DBPasseord string
		DBBase string
	}
	LogConfig struct{
		LogLevel string
	}
}

func (conf *Config) LoadFromEnv() error  {
	err := envconfig.Init(conf)
	return nil
}


func InitConfig() {
	
}

func InitContentConfig() error {

	viper.SetConfigName("config")

	viper.AddConfigPath("D:\\go\\goWork\\src\\ginTest\\config")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	viper.WatchConfig()

	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Info("Config file changed:", e.Name)
	})

	viper.AutomaticEnv()
	return nil
}
