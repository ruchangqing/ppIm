package utils

import (
	uuid "github.com/satori/go.uuid"
)

// 生成分布式uuid
func GenUuid() string {
	id := uuid.NewV4()
	return id.String()
}
