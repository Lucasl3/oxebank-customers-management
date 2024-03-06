package utils

import (
    "math/rand"
    "strconv"

    "customers-management/models"
    "customers-management/initializers"
)

func GenerateAccountNumber() string {
    accountNumber := rand.Intn(99999999-10000000+1) + 10000000
    return strconv.Itoa(accountNumber)
}

func IsAccountNumberExists(accountNumber string) bool {
	var user models.User
	if err := initializers.DB.Where("account_number = ?", accountNumber).First(&user).Error; err != nil {
		return false
	}
	return true
}
