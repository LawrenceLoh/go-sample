package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var (
	SysConfig Config
)

type Config struct {
	PostgreDriver  string `mapstructure:"POSTGRES_DRIVER"`
	PostgresSource string `mapstructure:"POSTGRES_SOURCE"`

	Port string `mapstructure:"PORT"`
}

func LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	err = viper.Unmarshal(&SysConfig)
	fmt.Println("Loading config successfully", SysConfig)
}
