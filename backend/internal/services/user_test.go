package services_test

import (
	"net/http"
	"testing"

	"github.com/baimhons/nom-naa-shop.git/internal/dtos/request"
	mock_repository "github.com/baimhons/nom-naa-shop.git/internal/mock/repository"
	mock_utils "github.com/baimhons/nom-naa-shop.git/internal/mock/utils"
	"github.com/baimhons/nom-naa-shop.git/internal/services"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type RegisterUserTestSuite struct {
	suite.Suite
	service      services.UserService
	mockRedis    *mock_utils.MockRedis
	mockBcrypt   *mock_utils.MockBcrypt
	mockUserRepo *mock_repository.MockUserRepository
	mockCartRepo *mock_repository.MockCartRepository
}

func (s *RegisterUserTestSuite) SetupTest() {
	s.mockUserRepo = new(mock_repository.MockUserRepository)
	s.mockCartRepo = new(mock_repository.MockCartRepository)
	s.mockRedis = new(mock_utils.MockRedis)
	s.mockBcrypt = new(mock_utils.MockBcrypt)
	s.service = services.NewUserService(s.mockUserRepo, s.mockCartRepo, s.mockRedis)
}

// func (s *RegisterUserTestSuite) TestRegisterUserSuccess() {
// 	// Mock the *gorm.DB to simulate transaction behavior
// 	mockTx := new(mock.Mock) // Mock *gorm.DB as transaction
// 	dbInstance := &gorm.DB{} // Create a real *gorm.DB instance

// 	// Mock Begin to return the mock transaction (Mock *gorm.DB instance)
// 	s.mockUserRepo.On("Begin").Return(mockTx)

// 	// Mock GetBy to return error for both username and email to check for non-existence
// 	s.mockUserRepo.On("GetBy", "username", "testuser", mock.AnythingOfType("*models.User")).Return(gorm.ErrRecordNotFound)
// 	s.mockUserRepo.On("GetBy", "email", "test@example.com", mock.AnythingOfType("*models.User")).Return(gorm.ErrRecordNotFound)

// 	// Mock Create User and Cart using mockTx (This simulates the transaction context)
// 	mockTx.On("Create", mock.AnythingOfType("*models.User")).Return(dbInstance)
// 	mockTx.On("Create", mock.AnythingOfType("*models.Cart")).Return(dbInstance)

// 	// Mock Commit and Rollback on the mock transaction (Simulating Commit and Rollback methods)
// 	mockTx.On("Commit").Return(dbInstance)
// 	mockTx.On("Rollback").Return(dbInstance)

// 	// Call RegisterUser
// 	_, status, err := s.service.RegisterUser(request.RegisterUser{
// 		Username:        "testuser",
// 		Email:           "test@example.com",
// 		Password:        "password123",
// 		ConfirmPassword: "password123",
// 		PhoneNumber:     "08123456",
// 		FirstName:       "John",
// 		LastName:        "Doe",
// 	})

// 	// Verify no error and status is 201 Created
// 	s.Require().NoError(err)
// 	s.Require().Equal(http.StatusCreated, status)

// 	// Assert mock calls were made correctly
// 	s.mockUserRepo.AssertCalled(s.T(), "Begin")
// 	s.mockUserRepo.AssertCalled(s.T(), "GetBy", "username", "testuser", mock.AnythingOfType("*models.User"))
// 	s.mockUserRepo.AssertCalled(s.T(), "GetBy", "email", "test@example.com", mock.AnythingOfType("*models.User"))
// 	mockTx.AssertCalled(s.T(), "Create", mock.AnythingOfType("*models.User"))
// 	mockTx.AssertCalled(s.T(), "Create", mock.AnythingOfType("*models.Cart"))
// 	mockTx.AssertCalled(s.T(), "Commit")
// }

func (s *RegisterUserTestSuite) TestRegisterUserExist() {
	s.mockUserRepo.On("GetBy", "email", "test@example.com", mock.Anything).Return(nil)

	_, status, err := s.service.RegisterUser(request.RegisterUser{
		Username:        "testuser",
		Email:           "test@example.com",
		Password:        "password123",
		ConfirmPassword: "password123",
		PhoneNumber:     "08123456",
		FirstName:       "John",
		LastName:        "Doe",
	})

	s.Require().Error(err)
	s.Require().Equal(http.StatusConflict, status)
	s.mockUserRepo.AssertCalled(s.T(), "GetBy", "email", "test@example.com", mock.Anything)
}

func TestRegisterUser(t *testing.T) {
	suite.Run(t, new(RegisterUserTestSuite))
}
