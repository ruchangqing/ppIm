package boot

import (
	"github.com/spf13/viper"
	_ "net/http/pprof"
)

func LoadConfig() {
	var filePath = "config.yml"
	viper.SetConfigType("yml")
	viper.SetConfigFile(filePath)
	err := viper.ReadInConfig()
	if err != nil {
		panic("读取配置文件出错：" + err.Error())
	}
}
