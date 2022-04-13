package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var c *viper.Viper

// Init initializes config
func Init(env string) {
	c = viper.New()
	c.SetConfigFile("yaml")
	c.SetConfigName(env)
	c.AddConfigPath("config/environments/")
	if err := c.ReadInConfig(); err != nil {
		fmt.Println(err.Error())
	}
}

// GetConfig returns config
func GetConfig() *viper.Viper {
	return c
}
