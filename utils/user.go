package utils

import (
	"customers-management/models"
	"customers-management/initializers"
)

func UserExists(userID uint) bool {
	var user models.User
	if err := initializers.DB.First(&user, userID).Error; err != nil {
		return false
	}
	return true
}