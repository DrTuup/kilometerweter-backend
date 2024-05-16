package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRegistrations(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "GET /registrations",
	})
}

func PostRegistrations(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "POST /registrations",
	})
}

func DeleteRegistration(c *gin.Context) {
	// Get the ID from the URL
	id := c.Param("id")

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "DELETE /registrations/" + id,
	})
}
