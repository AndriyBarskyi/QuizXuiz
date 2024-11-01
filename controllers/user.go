package controllers

import (
	"QuizXuiz/config"
	"QuizXuiz/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save user in DB
	if result := config.DB.Create(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully!"})
}

func LoginUser(c *gin.Context) {
	// Placeholder login functionality for now
	c.JSON(http.StatusOK, gin.H{"message": "Login not implemented yet"})
}
