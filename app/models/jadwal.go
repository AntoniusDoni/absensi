package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Jadwal struct {
	Id              uuid.UUID      `gorm:"primaryKey"`
	SchaduleDate    time.Time      `gorm:"type:date"`
	SchaduleTimeIn  datatypes.Time `gorm:"type:time"`
	SchaduleTimeOut datatypes.Time `gorm:"type:time"`
	CreatedAt       time.Time      `gorm:"autoCreateTime"`
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

func (jadwal *Jadwal) BeforeCreate(tx *gorm.DB) (err error) {
	jadwal.Id = uuid.New()
	return
}
