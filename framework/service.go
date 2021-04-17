package framework

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/olivere/elastic/v7"
	"github.com/spf13/viper"
	clientv3 "go.etcd.io/etcd/client/v3"
	"ppIm/global"
	"ppIm/model"
	"ppIm/servers"
	"time"
)

func LoadService() {
	connectDb()
	connectRedis()
	connectElasticsearch()
	connectEtcd()
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
		fmt.Println(err)
	}
	// 全局禁用表名复数
	global.Mysql.SingularTable(true)
}

func connectRedis() {
	host := viper.GetString("redis.host")
	port := viper.GetString("redis.port")
	pass := viper.GetString("redis.pass")
	global.Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: pass,
		DB:       0,
	})
	_, err := global.Redis.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println(err)
	}
}

func connectElasticsearch() {
	host := viper.GetString("elasticsearch.host")
	port := viper.GetString("elasticsearch.port")
	user := viper.GetString("elasticsearch.user")
	pass := viper.GetString("elasticsearch.pass")
	var err error
	global.Elasticsearch, err = elastic.NewClient(
		elastic.SetURL("http://"+host+":"+port),
		elastic.SetBasicAuth(user, pass),
	)
	if err != nil {
		fmt.Println(err)
	}

	// 创建用户位置索引
	_, _ = global.Elasticsearch.CreateIndex("user_location").BodyString(model.CreateUserLocationIndex).Do(context.Background())
}

func connectEtcd() {
	var err error
	servers.EtcdClient, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{viper.GetString("cluster.etcd")},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println(err)
	}
	//defer EtcdClient.Close()
}
