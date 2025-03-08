package repositories

import (
	"github.com/baimhons/nom-naa-shop.git/internal/dtos/request"
	"github.com/baimhons/nom-naa-shop.git/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReviewRepository interface {
	BaseRepository[models.Review]
	GetAllBySnackID(reviews *[]models.Review, querys *request.PaginationQuery, snackID uuid.UUID) error
}

type ReviewRepositoryImpl struct {
	BaseRepository[models.Review]
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) ReviewRepository {
	return &ReviewRepositoryImpl{
		BaseRepository: NewBaseRepository[models.Review](db),
		db:             db,
	}
}

func (r *ReviewRepositoryImpl) GetAllBySnackID(reviews *[]models.Review, querys *request.PaginationQuery, snackID uuid.UUID) error {
	return r.db.Where("snack_id = ?", snackID).Find(reviews).Error
}
