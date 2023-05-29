package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Jadwal struct {
	Id              uuid.UUID `gorm:"primaryKey"`
	SchaduleDate    time.Time
	SchaduleTimeIn  time.Time
	SchaduleTimeOut time.Time
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

func (jadwal *Jadwal) BeforeCreate(tx *gorm.DB) (err error) {
	jadwal.Id = uuid.New()
	return
}
