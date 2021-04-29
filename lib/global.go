package lib

import (
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	"github.com/olivere/elastic/v7"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
)

var (
	Db                  *gorm.DB                 //数据库客户端实例
	Redis               *redis.Client            //redis客户端实例
	Elasticsearch       *elastic.Client          //elasticsearch客户端实例
	JwtHmacSampleSecret = []byte("pancoiscool!") //jwt加密密钥
	ServerAddress       string                   //本机集群rpc地址
	Logger              *zap.SugaredLogger       //zap日志
	Etcd                *clientv3.Client         //etcd客户端
)
