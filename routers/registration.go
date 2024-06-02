package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"rubenclaessens.nl/kilometerweter-backend/handlers"
)

func getRegistrationsRouter(router *gin.Engine) {
	router.GET("/registrations", handlers.GetRegistrationsHandler)
}

func postRegistrationRouter(router *gin.Engine) {
	router.POST("/registrations", handlers.PostRegistrationHandler)
}

func deleteRegistrationRouter(router *gin.Engine) {
	router.DELETE("/registrations/:id", handlers.DeleteRegistrationHandler)
}

func updateRegistrationRouter(router *gin.Engine) {
	router.PUT("/registrations/:id", handlers.UpdateRegistrationHandler)
}

func SetupRouters() *gin.Engine {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "DELETE", "PUT"}

	router.Use(cors.New(config))

	getRegistrationsRouter(router)
	postRegistrationRouter(router)
	deleteRegistrationRouter(router)
	updateRegistrationRouter(router)

	return router
}
