package controllers

import (
	"customers-management/initializers"
	"customers-management/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CardsBlock(c *gin.Context) {
	id := c.Param("id")
	
	var card models.Card
	if err := initializers.DB.First(&card, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Card not found"})
		return
	}

	var body struct {
		Blocked *bool `json:"blocked" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := initializers.DB.Model(&card).Update("blocked", body.Blocked)

	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": card,
	})
}

func UsersBlock(c *gin.Context) {
	id := c.Param("id")
	
	var user models.User
	if err := initializers.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var body struct {
		AccountBlocked *bool `json:"account_blocked" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := initializers.DB.Model(&user).Update("account_blocked", body.AccountBlocked)

	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}