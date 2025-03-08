package repositories

import (
	"github.com/baimhons/nom-naa-shop.git/internal/models"
	"gorm.io/gorm"
)

type OrderRepository interface {
	BaseRepository[models.Order]
	GetDB() *gorm.DB
	GetOrder(order models.Order) (models.Order, error)
}

type OrderRepositoryImpl struct {
	BaseRepository[models.Order]
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &OrderRepositoryImpl{
		BaseRepository: NewBaseRepository[models.Order](db),
		db:             db,
	}
}

func (r *OrderRepositoryImpl) GetDB() *gorm.DB {
	return r.db
}

func (r *OrderRepositoryImpl) GetOrder(order models.Order) (models.Order, error) {
	return order, r.db.Where("tracking_id = ?", order.TrackingID).First(&order).Error
}
