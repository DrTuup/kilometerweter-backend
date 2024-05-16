package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"rubenclaessens.nl/kilometerweter-backend/models"
	"rubenclaessens.nl/kilometerweter-backend/routers"
)

func main() {
	// Connect and migrate database
	dsn := "host=localhost user=postgres password=postgres dbname=test port=5432 sslmode=disable TimeZone=Europe/Amsterdam"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.Registration{})

	routers.SetupRouters()
}
