package utils

import (
    "math/rand"
    "strconv"

    "customers-management/models"
    "customers-management/initializers"
)

func GenerateCardNumber() string {
		cardNumber := rand.Intn(9999999999999999-1000000000000000+1) + 1000000000000000
		return strconv.Itoa(cardNumber)
}

func IsCardNumberExists(cardNumber string) bool {
	var card models.Card
	if err := initializers.DB.Where("number = ?", cardNumber).First(&card).Error; err != nil {
		return false
	}
	return true
}