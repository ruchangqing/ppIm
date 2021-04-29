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
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"ppIm/global"
	"ppIm/model"
	"ppIm/utils"
	"strings"
	"time"
)

func LoadService() {
	InitLogger()
	connectDb()
	connectRedis()
	connectElasticsearch()
	connectEtcd()
	setQiNiu()
}

// 初始化日志
func InitLogger() {
	writeSyncer := GetLogWriter()
	encoder := GetEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller())
	global.Logger = logger.Sugar()
}

// 连接数据库
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

// 连接redis
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

// 连接es
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

// 连接etcd
func connectEtcd() {
	etcdString := viper.GetString("cluster.etcd")
	etcdArr := strings.Split(etcdString, "|")
	var err error
	global.Etcd, err = clientv3.New(clientv3.Config{
		Endpoints:   etcdArr,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic("连接Etcd出错：" + err.Error())
	}
	//defer global.Etcd.Close()
}

// 设置七牛云对象存储
func setQiNiu() {
	utils.QiNiuClient.AccessKey = viper.GetString("qiniu.accessKey")
	utils.QiNiuClient.SecretKey = viper.GetString("qiniu.secretKey")
	utils.QiNiuClient.Bucket = viper.GetString("qiniu.bucket")
	utils.QiNiuClient.Domain = viper.GetString("qiniu.domain")
}