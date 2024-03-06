package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Email string
	Password string
	AccountNumber string
	AgencyNumber string
	AccountBlocked bool
	AcceptedTerms bool
	Cards []Card `gorm:"foreignKey:UserID"`
}