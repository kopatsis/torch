package data

import (
	"torch/data/repositories"
	"torch/data/services"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type MainService struct {
	User    services.UserService
	Product services.ProductService
}

func NewMainService(db *gorm.DB, redis *redis.Client) *MainService {
	return &MainService{
		User:    services.NewUserService(repositories.NewUserRepository(db, redis)),
		Product: services.NewProductService(repositories.NewProductRepository(db, redis)),
	}
}
