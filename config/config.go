package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Configurations struct {
	Hue HueConfig
}

type HueConfig struct {
	BaseUrl string
	Colors map[string]int
}

func readConfigFile() *Configurations {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("error reading config file %v\n", err)
	}

	var config Configurations

	err := viper.Unmarshal(&config)
	if err != nil {
		fmt.Printf("unable to decode into struct %v\n", err)
	}

	return &config
}

// GetConfig - creates a new HueConfig
func GetConfig() *Configurations {
	config := readConfigFile()
	return config;
}