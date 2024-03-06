package models

import (
	"gorm.io/gorm"
)

type Card struct {
  gorm.Model
  Number string
	Password string
	Blocked bool
  UserID uint
}