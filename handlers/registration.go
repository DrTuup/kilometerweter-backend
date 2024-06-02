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

func UpdateRegistrationHandler(c *gin.Context) {
	id := c.Param("id")

	// check if the registration exists
	if !models.RegistrationExists(id) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Registration with id " + id + " not found"})
		return
	}

	// bind the request body to the registration struct
	var registration models.Registration
	if err := c.ShouldBindJSON(&registration); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update the registration
	err := models.UpdateRegistration(id, registration)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update registration"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Updated registration with id " + id,
	})
}
