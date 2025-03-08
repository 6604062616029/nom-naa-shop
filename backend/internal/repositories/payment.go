package repositories

import (
	"github.com/baimhons/nom-naa-shop.git/internal/models"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	BaseRepository[models.Payment]
	GetDB() *gorm.DB
}

type PaymentRepositoryImpl struct {
	BaseRepository[models.Payment]
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &PaymentRepositoryImpl{
		BaseRepository: NewBaseRepository[models.Payment](db),
		db:             db,
	}
}

func (r *PaymentRepositoryImpl) GetDB() *gorm.DB {
	return r.db
}
