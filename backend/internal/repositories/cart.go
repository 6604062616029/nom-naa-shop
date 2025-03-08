package repositories

import (
	"github.com/baimhons/nom-naa-shop.git/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CartRepository interface {
	BaseRepository[models.Cart]
	GetCartByCondition(condition string, args ...interface{}) (cart *models.Cart, err error)
	GetCartByID(id uuid.UUID) (*models.Cart, error)
	GetDB() *gorm.DB
	GetCartWithItems(cartID uuid.UUID) (*models.Cart, error)
	GetCartWithItemsAndSnack(cartID uuid.UUID) (*models.Cart, error)
}

type CartRepositoryImpl struct {
	BaseRepository[models.Cart]
	DB *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepositoryImpl {
	return &CartRepositoryImpl{
		BaseRepository: NewBaseRepository[models.Cart](db),
		DB:             db,
	}
}

func (r *CartRepositoryImpl) GetCartByCondition(condition string, args ...interface{}) (cart *models.Cart, err error) {
	return cart, r.DB.Where(condition, args...).Preload("Items").Find(&cart).Error
}

func (r *CartRepositoryImpl) GetCartByID(id uuid.UUID) (*models.Cart, error) {
	var cart models.Cart
	if err := r.DB.Preload("Items.Snack").Where("id = ?", id).First(&cart).Error; err != nil {
		return nil, err
	}
	return &cart, nil
}

func (r *CartRepositoryImpl) GetDB() *gorm.DB {
	return r.DB
}

func (r *CartRepositoryImpl) GetCartWithItems(cartID uuid.UUID) (*models.Cart, error) {
	var cart models.Cart
	if err := r.DB.
		Preload("Items").
		Preload("Items.Snack").
		First(&cart, cartID).Error; err != nil {
		return nil, err
	}
	return &cart, nil
}

func (r *CartRepositoryImpl) GetCartWithItemsAndSnack(cartID uuid.UUID) (*models.Cart, error) {
	var cart models.Cart
	if err := r.DB.
		Preload("Items", func(db *gorm.DB) *gorm.DB {
			return db.Order("items.id")
		}).
		Preload("Items.Snack").
		Where("id = ?", cartID).
		First(&cart).Error; err != nil {
		return nil, err
	}
	return &cart, nil
}
