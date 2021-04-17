package framework

import (
	"fmt"
	"github.com/spf13/viper"
	"ppIm/global"
)

func LoadConfig() {
	var filePath = "config.yml"
	viper.SetConfigType("yml")
	viper.SetConfigFile(filePath)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	global.IsCluster = viper.GetBool("cluster.open")
}