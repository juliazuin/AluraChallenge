package config

import (
	"log"

	"github.com/go-playground/validator"
	"github.com/spf13/viper"
)

type Config struct {
	Viper    *viper.Viper
	Validate *validator.Validate
}

var validate *validator.Validate

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

	validate = validator.New()

	return &Config{
		Viper:    viper.GetViper(),
		Validate: validate,
	}

}
