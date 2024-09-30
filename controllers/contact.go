package controllers

import (
	"fmt"
	"net/http"
	"regexp"
	"tsa-innovation-lab/database"
	"tsa-innovation-lab/models"

	"github.com/gin-gonic/gin"
)

func CreateContact(c *gin.Context) {
	var contact models.Contact

	if err := c.BindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	if contact.Email == "" || !isValidEmail(contact.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email"})
		return
	}

	formattedPhone, err := formatPhoneNumber(contact.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid phone number"})
		return
	}
	contact.PhoneNumber = formattedPhone

	result := database.DB.Create(&contact)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save contact"})
		return
	}

	c.JSON(http.StatusOK, contact)
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func formatPhoneNumber(phone string) (string, error) {
	re := regexp.MustCompile(`^\+61\d{9}$`)
	if re.MatchString(phone) {
		return phone, nil
	}
	return "", fmt.Errorf("phone number is not in valid E.164 format")
}
