package models

import (
	"gorm.io/gorm"
)

type Term struct {
	gorm.Model
	Text string
	Version string
}