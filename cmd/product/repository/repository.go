package repository

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ProductRepository struct {
	Redis    *redis.Client
	Database *gorm.DB
}

func NewProductRepository(redis *redis.Client, db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		Redis:    redis,
		Database: db,
	}
}
