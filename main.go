package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"rubenclaessens.nl/kilometerweter-backend/config"
	"rubenclaessens.nl/kilometerweter-backend/models"
	"rubenclaessens.nl/kilometerweter-backend/routers"
)

func main() {
	// Connect and migrate database
	config.DB, _ = gorm.Open(postgres.Open("host=localhost user=postgres password=postgres dbname=kilometerweter port=5432"), &gorm.Config{})
	config.DB.AutoMigrate(&models.Registration{})

	r := routers.SetupRouters()

	r.Run(":8080")
}
