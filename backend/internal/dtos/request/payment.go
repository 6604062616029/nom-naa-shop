package request

import (
	"fmt"
	"io"
	"mime/multipart"

	"github.com/baimhons/nom-naa-shop.git/internal/models"
	"github.com/google/uuid"
)

type PaymentRequest struct {
	OrderID uuid.UUID               `form:"order_id" validate:"required"`
	Proof   []*multipart.FileHeader `form:"files"`
}

func (r *PaymentRequest) ToModel() (*models.Payment, error) {
	if len(r.Proof) == 0 {
		return nil, fmt.Errorf("no proof file uploaded")
	}

	// Open the proof image file
	proofFile, err := r.Proof[0].Open()
	if err != nil {
		return nil, err
	}
	defer proofFile.Close()

	// Read the file content
	proofBytes, err := io.ReadAll(proofFile)
	if err != nil {
		return nil, err
	}

	return &models.Payment{
		OrderID: r.OrderID,
		Proof:   proofBytes,
	}, nil
}
