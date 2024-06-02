package models

import (
	"time"

	"rubenclaessens.nl/kilometerweter-backend/config"
)

type Registration struct {
	ID          uint      `json:"id" gorm:"primaryKey,autoIncrement"`
	Description string    `json:"description"`
	Kilometers  uint      `json:"kilometers"`
	Date        time.Time `json:"date"`
}

func GetRegistrations() []Registration {
	registrations := []Registration{}
	config.DB.Find(&registrations)
	return registrations
}

func CreateRegistration(r Registration) (err error) {
	if err = config.DB.Create(&r).Error; err != nil {
		return err
	}
	return nil
}

func DeleteRegistration(id string) (err error) {
	if err = config.DB.Delete(&Registration{}, id).Error; err != nil {
		return err
	}
	return nil
}

func UpdateRegistration(id string, r Registration) (err error) {
	if err = config.DB.Model(&Registration{}).Where("id = ?", id).Updates(r).Error; err != nil {
		return err
	}
	return nil
}

func RegistrationExists(id string) bool {
	var registration Registration
	config.DB.Where("id = ?", id).First(&registration)
	return registration.ID != 0
}
