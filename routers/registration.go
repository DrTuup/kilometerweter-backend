package routers

import (
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

func SetupRouters() *gin.Engine {
	router := gin.Default()

	getRegistrationsRouter(router)
	postRegistrationRouter(router)
	deleteRegistrationRouter(router)

	return router
}
