package config

import (
	"log"

	"github.com/spf13/viper"
)

type DatabaseConfigurations struct {
	DBName     string
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
}

type Config struct {
	Viper *viper.Viper
}

func NewConfig() *Config {
	viper.AddConfigPath(".")
	// viper.SetConfigName("env")
	viper.SetConfigFile("env.yaml")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Print(err.Error())
	} else {
		log.Print("Config loaded...")
	}

	return &Config{
		Viper: viper.GetViper(),
	}

}
