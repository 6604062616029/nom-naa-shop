package repositories

import (
	"github.com/baimhons/nom-naa-shop.git/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SnackRepository interface {
	BaseRepository[models.Snack]
	GetSnackByID(id uuid.UUID) (snack *models.Snack, err error)
	GetSnackByIDWithReview(id uuid.UUID) (snack *models.Snack, err error)
}

type snackRepositoryImpl struct {
	BaseRepository[models.Snack]
	DB *gorm.DB
}

func NewSnackRepository(db *gorm.DB) SnackRepository {
	return &snackRepositoryImpl{
		BaseRepository: NewBaseRepository[models.Snack](db),
		DB:             db,
	}
}

func (r *snackRepositoryImpl) GetSnackByID(id uuid.UUID) (snack *models.Snack, err error) {
	return snack, r.DB.Where("id = ?", id).First(&snack).Error
}

func (r *snackRepositoryImpl) GetSnackByIDWithReview(id uuid.UUID) (snack *models.Snack, err error) {
	return snack, r.DB.Where("id = ?", id).Preload("Reviews").First(&snack).Error
}
