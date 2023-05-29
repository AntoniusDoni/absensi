package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Settings struct {
	Id             uuid.UUID `gorm:"primaryKey"`
	Name           string
	Longitude      string
	Latitute       string
	EnableLocation bool
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}
