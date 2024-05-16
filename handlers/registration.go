package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"rubenclaessens.nl/kilometerweter-backend/models"
)

func GetRegistrationsHandler(c *gin.Context) {
	registrations := models.GetRegistrations()
	c.ShouldBindJSON(&registrations)

	c.IndentedJSON(http.StatusOK, gin.H{
		"registrations": registrations,
	})
}

func PostRegistrationHandler(c *gin.Context) {
	// bind the request body to the registration struct
	var registration models.Registration
	if err := c.ShouldBindJSON(&registration); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create the registration
	err := models.CreateRegistration(registration)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create registration"})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"registration": registration})
}

func DeleteRegistrationHandler(c *gin.Context) {
	id := c.Param("id")

	// delete the registration
	err := models.DeleteRegistration(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete registration"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Deleted registration with id " + id,
	})
}
