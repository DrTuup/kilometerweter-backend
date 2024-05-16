package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"rubenclaessens.nl/kilometerweter-backend/config"
	"rubenclaessens.nl/kilometerweter-backend/models"
	"rubenclaessens.nl/kilometerweter-backend/routers"
)

func main() {
	godotenv.Load()
	DB_HOST := os.Getenv("DB_HOST")
	DB_USERNAME := os.Getenv("DB_USERNAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT := os.Getenv("DB_PORT")
	APP_PORT := os.Getenv("APP_PORT")
	GIN_MODE := os.Getenv("GIN_MODE")

	if GIN_MODE == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Connect and migrate database
	config.DB, _ = gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", DB_HOST, DB_USERNAME, DB_PASSWORD, DB_NAME, DB_PORT)), &gorm.Config{})
	config.DB.AutoMigrate(&models.Registration{})

	r := routers.SetupRouters()

	r.Run(":" + APP_PORT)

}
