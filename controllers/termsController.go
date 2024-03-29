package controllers

import (
	"customers-management/initializers"
	"customers-management/models"
	"net/http"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"errors"
)

func TermsCreate(c *gin.Context) {
	var body struct {
		Text string `json:"text" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var lastTerm models.Term
	if err := initializers.DB.Order("version DESC").First(&lastTerm).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching the latest version"})
			return
		}
	}

	newVersion := 1
	if lastTerm.Version != 0 {
		newVersion = lastTerm.Version + 1
	}

	term := models.Term{
		Text: body.Text,
		Version: newVersion,
	}

	result := initializers.DB.Create(&term)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, gin.H{
		"data": term,
	})
}

func TermsIndex(c *gin.Context) {
	var terms []models.Term
	initializers.DB.Find(&terms)

	c.JSON(200, gin.H{
		"data": terms,
	})
}

func TermsShow(c *gin.Context) {
	version := c.Param("version")

	var term models.Term
	if err := initializers.DB.Where("version = ?", version).First(&term).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Version term not found"})
		return
	}

	c.JSON(200, gin.H{
		"data": term,
	})
}

func TermsLatest(c *gin.Context) {
	var term models.Term
	if err := initializers.DB.Order("version DESC").First(&term).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No terms found"})
		return
	}

	c.JSON(200, gin.H{
		"data": term,
	})
}

func TermsAccept(c *gin.Context) {
	var body struct {
		UserID uint `json:"user_id" binding:"required"`
		Version int `json:"version" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := initializers.DB.First(&user, body.UserID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	var term models.Term
	if err := initializers.DB.Where("version = ?", body.Version).First(&term).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Version term not found"})
		return
	}

	user.AcceptedTerms = true
	user.AccountBlocked = false
	result := initializers.DB.Save(&user)

	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(200, gin.H{
		"message": "Terms accepted",
		"data": user,
	})
}