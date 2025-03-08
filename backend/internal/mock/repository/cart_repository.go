package mock_repository

import (
	"github.com/baimhons/nom-naa-shop.git/internal/dtos/request"
	"github.com/baimhons/nom-naa-shop.git/internal/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockCartRepository struct {
	mock.Mock
}

func (m *MockCartRepository) GetCartByCondition(condition string, args ...interface{}) (cart *models.Cart, err error) {
	return m.Called(condition, args).Get(0).(*models.Cart), m.Called(condition, args).Error(0)
}

func (m *MockCartRepository) GetCartByID(id uuid.UUID) (*models.Cart, error) {
	return m.Called(id).Get(0).(*models.Cart), m.Called(id).Error(0)
}

func (m *MockCartRepository) GetByID(item *models.Cart, id uuid.UUID) error {
	return m.Called(item, id).Error(0)
}

func (m *MockCartRepository) GetAll(items *[]models.Cart, pagination *request.PaginationQuery) error {
	return m.Called(items, pagination).Error(0)
}

func (m *MockCartRepository) GetBy(field string, value string, item *models.Cart) error {
	return m.Called(field, value, item).Error(0)
}

func (m *MockCartRepository) CreateMany(items *[]models.Cart) error {
	return m.Called(items).Error(0)
}

func (m *MockCartRepository) Update(item *models.Cart) error {
	return m.Called(item).Error(0)
}

func (m *MockCartRepository) Create(item *models.Cart) error {
	return m.Called(item).Error(0)
}

func (m *MockCartRepository) Delete(item *models.Cart) error {
	return m.Called(item).Error(0)
}

func (m *MockCartRepository) Begin() *gorm.DB {
	return m.Called().Get(0).(*gorm.DB)
}

func (m *MockCartRepository) GetDB() *gorm.DB {
	return m.Called().Get(0).(*gorm.DB)
}

func (m *MockCartRepository) GetCartWithItems(cartID uuid.UUID) (*models.Cart, error) {
	return m.Called(cartID).Get(0).(*models.Cart), m.Called(cartID).Error(0)
}

func (m *MockCartRepository) GetCartWithItemsAndSnack(cartID uuid.UUID) (*models.Cart, error) {
	return m.Called(cartID).Get(0).(*models.Cart), m.Called(cartID).Error(0)
}
