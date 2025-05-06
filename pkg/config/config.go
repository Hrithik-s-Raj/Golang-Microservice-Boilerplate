package config

import (
	"log"

	"github.com/spf13/viper"
)

func LoadEnv(env string) {
	println(env)
	viper.SetConfigName(env)
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	viper.AutomaticEnv()
}

func GetPort() string {
	if port := viper.GetString("PORT"); port != "" {
		return port
	}
	return "8080"
}

func GetMongoURI() string {
	return viper.GetString("MONGODB_URI")
}
