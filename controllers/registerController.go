package controllers

import (
	"customers-management/initializers"
	"customers-management/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TermsCreate(c *gin.Context) {
	var body struct {
		Text string `json:"text" binding:"required"`
		Version string `json:"version" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	term := models.Term{
		Text: body.Text,
		Version: body.Version,
	}
	result := initializers.DB.Create(&term)

	if result.Error != nil {
		c.Status(http.StatusBadRequest)
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

func TermsAccept(c *gin.Context) {
	var body struct {
		UserID uint `json:"user_id" binding:"required"`
		Version string `json:"version" binding:"required"`
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