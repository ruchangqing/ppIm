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
	"ppIm/services"
	"strings"
	"time"
)

func LoadService() {
	connectDb()
	connectRedis()
	connectElasticsearch()
	connectEtcd()
	setQiNiu()
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
	global.Db, err = gorm.Open(dbType, args)
	if err != nil {
		panic("连接数据库出错：" + err.Error())
	}
	// 全局禁用表名复数
	global.Db.SingularTable(true)
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
		panic("连接Redis出错：" + err.Error())
	}
}

func connectElasticsearch() {
	host := viper.GetString("elasticsearch.host")
	port := viper.GetString("elasticsearch.port")
	user := viper.GetString("elasticsearch.user")
	pass := viper.GetString("elasticsearch.pass")
	var err error
	global.Elasticsearch, err = elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL("http://"+host+":"+port),
		elastic.SetBasicAuth(user, pass),
	)
	if err != nil {
		panic("连接Elasticsearch出错：" + err.Error())
	}

	// 创建用户位置索引
	_, _ = global.Elasticsearch.CreateIndex("user_location").BodyString(model.CreateUserLocationIndex).Do(context.Background())
}

func connectEtcd() {
	etcdString := viper.GetString("cluster.etcd")
	etcdArr := strings.Split(etcdString, "|")
	var err error
	servers.EtcdClient, err = clientv3.New(clientv3.Config{
		Endpoints:   etcdArr,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic("连接Etcd出错：" + err.Error())
	}
	//defer EtcdClient.Close()
}

func setQiNiu() {
	services.QiNiuClient.AccessKey = viper.GetString("qiniu.accessKey")
	services.QiNiuClient.SecretKey = viper.GetString("qiniu.secretKey")
	services.QiNiuClient.Bucket = viper.GetString("qiniu.bucket")
	services.QiNiuClient.Domain = viper.GetString("qiniu.domain")
}
