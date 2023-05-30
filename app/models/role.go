package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Roles struct {
	Id        uuid.UUID      `gorm:"primaryKey" json:"roleId,omitempty"`
	Name      string         `json:"roles_name," `
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"create_at,omitempty"`
	UpdatedAt time.Time      `json:"update_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:",omitempty"`
}

func (role *Roles) BeforeCreate(tx *gorm.DB) (err error) {
	role.Id = uuid.New()
	return
}
