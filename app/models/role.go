package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Roles struct {
	Id        uuid.UUID `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (role *Roles) BeforeCreate(tx *gorm.DB) (err error) {
	role.Id = uuid.New()
	return
}
