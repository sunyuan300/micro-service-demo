package data

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)
type Data struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewData() (*Data,error) {
	return &Data{},nil
}