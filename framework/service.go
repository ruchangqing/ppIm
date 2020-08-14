package framework

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"ppIm/global"
)

func LoadService() {
	connectDb()
	connectRedis()
}

func connectDb() {
	dbType := viper.GetString("db.dbType")
	host := viper.GetString("db.host")
	port := viper.GetString("db.port")
	user := viper.GetString("db.user")
	pass := viper.GetString("db.pass")
	dbname := viper.GetString("db.dbname")
	charset := viper.GetString("db.charset")
	var err error
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", user, pass, host, port, dbname, charset)
	global.Mysql, err = gorm.Open(dbType, args)
	if err != nil {
		panic(err)
	}
	// 全局禁用表名复数
	global.Mysql.SingularTable(true)
}

func connectRedis() {
	host := viper.GetString("redis.host")
	port := viper.GetString("redis.port")
	pass := viper.GetString("redis.pass")
	global.RedisCtx = context.Background()
	global.Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: pass,
		DB:       0,
	})
	_, err := global.Redis.Ping(global.RedisCtx).Result()
	if err != nil {
		panic(err)
	}
}
