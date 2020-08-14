package global

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

var (
	Mysql               *gorm.DB
	Redis               *redis.Client
	RedisCtx            context.Context
	JwtHmacSampleSecret = []byte("pancoiscool!") // jwt加密密钥
)
