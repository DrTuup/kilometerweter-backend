package routers

import (
	"github.com/gin-gonic/gin"
	"rubenclaessens.nl/kilometerweter-backend/handlers"
)

func GetRegistrationsRouter(router *gin.Engine) {
	router.GET("/registrations", handlers.GetRegistrations)
}

func PostRegistrationsRouter(router *gin.Engine) {
	router.POST("/registrations", handlers.PostRegistrations)
}

func DeleteRegistrationRouter(router *gin.Engine) {
	router.DELETE("/registrations/:id", handlers.DeleteRegistration)
}

func SetupRouters() {
	router := gin.Default()

	GetRegistrationsRouter(router)
	PostRegistrationsRouter(router)
	DeleteRegistrationRouter(router)

	router.Run(":8080")
}
