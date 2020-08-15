package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	"github.com/olivere/elastic/v7"
)

var (
	Mysql               *gorm.DB
	Redis               *redis.Client
	Elasticsearch		*elastic.Client
	JwtHmacSampleSecret = []byte("pancoiscool!") // jwt加密密钥
)
