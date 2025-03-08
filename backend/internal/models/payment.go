package models

import (
	"github.com/google/uuid"
)

type Payment struct {
	BaseModel
	OrderID uuid.UUID `gorm:"not null"`
	Order   Order     `gorm:"foreignKey:OrderID"`
	Amount  float64   `gorm:"not null"`
	Proof   []byte    `gorm:"type:bytea"`
}
