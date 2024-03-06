package main

import (
	"customers-management/initializers"
	"customers-management/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
	initializers.DB.AutoMigrate(&models.Card{})
	initializers.DB.AutoMigrate(&models.Term{})
}