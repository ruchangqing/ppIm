package framework

import (
	"github.com/spf13/viper"
)

func LoadConfig() {
	var filePath = "config.yml"
	viper.SetConfigType("yml")
	viper.SetConfigFile(filePath)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}