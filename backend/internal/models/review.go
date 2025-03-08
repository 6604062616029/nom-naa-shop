package models

import "github.com/google/uuid"

type Review struct {
	BaseModel
	UserID  uuid.UUID `gorm:"not null"`
	User    User      `gorm:"foreignKey:UserID"`
	SnackID uuid.UUID `gorm:"not null"`
	Snack   Snack     `gorm:"foreignKey:SnackID"`
	Rating  int       `gorm:"rating;not null"`
	Comment string    `gorm:"comment;not null"`
}
