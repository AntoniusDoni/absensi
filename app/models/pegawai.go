package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Pegawai struct {
	Id        uuid.UUID `gorm:"primaryKey"`
	IdRole    uuid.UUID
	Nip       string `gorm:"index"`
	Name      string `gorm:"index"`
	Password  string
	Phone     string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (pegawai *Pegawai) BeforeCreate(tx *gorm.DB) (err error) {
	pegawai.Id = uuid.New()
	return
}
