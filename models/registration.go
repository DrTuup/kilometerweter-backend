package models

import (
	"time"
)

type Registration struct {
	ID          uint `gorm:"primaryKey,autoIncrement"`
	Description string
	Kilometers  uint
	Date        time.Time
}
