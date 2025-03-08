package services

import (
	"fmt"
	"io"
	"net/http"

	"github.com/baimhons/nom-naa-shop.git/internal/dtos/request"
	"github.com/baimhons/nom-naa-shop.git/internal/dtos/response"
	"github.com/baimhons/nom-naa-shop.git/internal/models"
	"github.com/baimhons/nom-naa-shop.git/internal/repositories"
	"github.com/google/uuid"
)

type SnackService interface {
	CreateSnack(req request.CreateSnackRequest) (resp response.SuccessResponse, statusCode int, err error)
	GetAllSnacks(querys request.PaginationQuery) (resp response.SuccessResponse, statusCode int, err error)
	GetSnackByID(id uuid.UUID) (resp response.SuccessResponse, statusCode int, err error)
	UpdateSnack(req request.UpdateSnackRequest, userContext models.UserContext, id uuid.UUID) (resp response.SuccessResponse, statusCode int, err error)
	DeleteSnack(id uuid.UUID) (resp response.SuccessResponse, statusCode int, err error)
	CreateReview(req request.CreateReviewRequest, userContext models.UserContext) (resp response.SuccessResponse, statusCode int, err error)
	GetAllReviewsBySnackID(querys request.PaginationQuery, snackID uuid.UUID) (resp response.SuccessResponse, statusCode int, err error)
}

type snackServiceImpl struct {
	snackRepository  repositories.SnackRepository
	reviewRepository repositories.ReviewRepository
}

func NewSnackService(snackRepository repositories.SnackRepository, reviewRepository repositories.ReviewRepository) SnackService {
	return &snackServiceImpl{snackRepository: snackRepository, reviewRepository: reviewRepository}
}

func (s *snackServiceImpl) CreateSnack(req request.CreateSnackRequest) (resp response.SuccessResponse, statusCode int, err error) {
	snack, err := req.ToModel()
	if err != nil {
		return resp, http.StatusInternalServerError, err
	}

	if err := s.snackRepository.Create(snack); err != nil {
		return resp, http.StatusInternalServerError, err
	}

	return response.SuccessResponse{
		Message: "Snack created successfully",
		Data:    snack,
	}, http.StatusCreated, nil
}

func (s *snackServiceImpl) GetAllSnacks(querys request.PaginationQuery) (resp response.SuccessResponse, statusCode int, err error) {
	snacks := []models.Snack{}
	if err := s.snackRepository.GetAll(&snacks, &querys); err != nil {
		return resp, http.StatusInternalServerError, err
	}

	return response.SuccessResponse{
		Message: "Snacks fetched successfully",
		Data:    snacks,
	}, http.StatusOK, nil
}

func (s *snackServiceImpl) GetSnackByID(id uuid.UUID) (resp response.SuccessResponse, statusCode int, err error) {
	snack := models.Snack{}
	if err := s.snackRepository.GetByID(&snack, id); err != nil {
		return resp, http.StatusInternalServerError, err
	}

	return response.SuccessResponse{
		Message: "Snack fetched successfully",
		Data:    snack,
	}, http.StatusOK, nil
}

func (s *snackServiceImpl) UpdateSnack(req request.UpdateSnackRequest, userContext models.UserContext, id uuid.UUID) (resp response.SuccessResponse, statusCode int, err error) {
	if userContext.ID == "" {
		return resp, http.StatusUnauthorized, fmt.Errorf("unauthorized user")
	}

	var existingSnack models.Snack
	if err := s.snackRepository.GetByID(&existingSnack, id); err != nil {
		return resp, http.StatusNotFound, fmt.Errorf("snack not found")
	}

	if req.Price != 0 {
		existingSnack.Price = req.Price
	}
	if req.Quantity != 0 {
		existingSnack.Quantity = req.Quantity
	}
	if len(req.Files) > 0 {
		imageFile, err := req.Files[0].Open()
		if err != nil {
			return resp, http.StatusInternalServerError, err
		}
		defer imageFile.Close()

		imageBytes, err := io.ReadAll(imageFile)
		if err != nil {
			return resp, http.StatusInternalServerError, err
		}
		existingSnack.Image = imageBytes
	}

	if err := s.snackRepository.Update(&existingSnack); err != nil {
		return resp, http.StatusInternalServerError, err
	}

	return response.SuccessResponse{
		Message: "Snack updated successfully",
		Data:    existingSnack,
	}, http.StatusOK, nil
}

func (s *snackServiceImpl) CreateReview(req request.CreateReviewRequest, userContext models.UserContext) (resp response.SuccessResponse, statusCode int, err error) {
	if userContext.ID == "" {
		return resp, http.StatusUnauthorized, fmt.Errorf("unauthorized user")
	}

	userID, err := uuid.Parse(userContext.ID)
	if err != nil {
		return resp, http.StatusInternalServerError, err
	}

	review := models.Review{
		UserID:  userID,
		SnackID: req.SnackID,
		Rating:  req.Rating,
		Comment: req.Comment,
	}

	if err := s.reviewRepository.Create(&review); err != nil {
		return resp, http.StatusInternalServerError, err
	}

	return response.SuccessResponse{
		Message: "Review created successfully",
		Data:    review,
	}, http.StatusCreated, nil
}

func (s *snackServiceImpl) GetAllReviewsBySnackID(querys request.PaginationQuery, snackID uuid.UUID) (resp response.SuccessResponse, statusCode int, err error) {
	reviews := []models.Review{}
	if err := s.reviewRepository.GetAllBySnackID(&reviews, &querys, snackID); err != nil {
		return resp, http.StatusInternalServerError, err
	}

	return response.SuccessResponse{
		Message: "Reviews fetched successfully",
		Data:    reviews,
	}, http.StatusOK, nil

}

func (s *snackServiceImpl) DeleteSnack(id uuid.UUID) (resp response.SuccessResponse, statusCode int, err error) {
	if err := s.snackRepository.Delete(&models.Snack{BaseModel: models.BaseModel{ID: id}}); err != nil {
		return resp, http.StatusInternalServerError, err
	}
	return response.SuccessResponse{
		Message: "Snack deleted successfully",
	}, http.StatusOK, nil
}
