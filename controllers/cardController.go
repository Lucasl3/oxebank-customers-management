package controllers

import (
	"math/rand"
	
	"customers-management/initializers"
	"customers-management/models"
	"customers-management/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CardsCreate(c *gin.Context) {
	var body struct {
		UserID uint `json:"user_id" binding:"required"`
	}
	
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verifica se o usuário com o ID fornecido existe
	if !utils.UserExists(body.UserID) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User with the provided ID does not exist"})
		return
	}

	// Gera um número de cartão aleatório
	cardNumber := utils.GenerateCardNumber()
	
	// Verifica se o número de cartão gerado já existe
	for utils.IsCardNumberExists(cardNumber) {
		cardNumber = utils.GenerateCardNumber()
	}
	
	var user models.User
	initializers.DB.First(&user, body.UserID)

	if !user.AcceptedTerms {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User has not accepted the terms"})
		return
	
	}

	if user.AccountBlocked {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User account is blocked"})
		return
	}

	card := models.Card{
		Number: cardNumber,
		Password: strconv.Itoa(rand.Intn(9999-1000+1) + 1000),
		UserID: body.UserID,
		Blocked: false,
	}
	result := initializers.DB.Create(&card)
	
	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(200, gin.H{
		"card": card,
	})
}

func CardsIndex(c *gin.Context) {
	var cards []models.Card
	initializers.DB.Preload("User").Find(&cards)

	c.JSON(200, gin.H{
		"data": cards,
	})
}

func CardsShow(c *gin.Context) {
	id := c.Param("id")
	
	var card models.Card
	if err := initializers.DB.First(&card, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Card not found"})
			return
	}

	c.JSON(http.StatusOK, gin.H{
			"data": card,
	})
}

func CardsIndexByUser(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
	}
	// Verifica se o usuário com o ID fornecido existe
	if !utils.UserExists(uint(userID)) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User with the provided ID does not exist"})
		return
	}

	var cards []models.Card
	if err := initializers.DB.Where("user_id = ?", userID).Find(&cards).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No cards found for the user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": cards,
	})
}

func CardsDelete(c *gin.Context) {
	id := c.Param("id")
	
	var card models.Card
	if err := initializers.DB.First(&card, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Card not found"})
		return
	}

	result := initializers.DB.Delete(&card)

	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Card deleted successfully",
	})
}
