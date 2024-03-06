package controllers

import (
	"customers-management/utils"
	"customers-management/initializers"
	"customers-management/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UsersCreate(c *gin.Context) {
	var body struct {
		Name string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Gera um número de conta aleatório
	accountNumber := utils.GenerateAccountNumber()

	// Verifica se o número de conta gerado já existe
	for utils.IsAccountNumberExists(accountNumber) {
		accountNumber = utils.GenerateAccountNumber()
	}

	user := models.User{
		Name: body.Name,
		Email: body.Email,
		Password: body.Password,
		AccountNumber: accountNumber,
		AgencyNumber: "0001",
		AccountBlocked: true,
		AcceptedTerms: false,
	}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}

func UsersIndex(c *gin.Context) {
	var users []models.User
	initializers.DB.Preload("Cards").Find(&users)

	c.JSON(200, gin.H{
		"data": users,
	})
}

func UsersShow(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := initializers.DB.Preload("Cards").First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, gin.H{
		"data": user,
	})
}

func UsersUpdate(c *gin.Context) {
	id := c.Param("id")
	
	var body struct {
		Name string `json:"name"`
		Email string `json:"email"`
		Password string `json:"password"`
		AccountBlocked bool `json:"account_blocked"`
	}
	
	c.Bind(&body)

	var user models.User
	initializers.DB.First(&user, id)

	initializers.DB.Model(&user).Updates(models.User{
		Name: body.Name,
		Email: body.Email,
		Password: body.Password,
		AccountBlocked: body.AccountBlocked,
	})

	c.JSON(200, gin.H{
		"data": user,
	})
}

func UsersDelete(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	initializers.DB.Delete(&user, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}