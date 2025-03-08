package repositories

import (
	"github.com/baimhons/nom-naa-shop.git/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	BaseRepository[models.User]
}

type userRepositoryImpl struct {
	BaseRepository[models.User]
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		BaseRepository: NewBaseRepository[models.User](db),
		DB:             db,
	}
}
