package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Absesnsi struct {
	Id        uuid.UUID `gorm:"primaryKey"`
	IdPegawai uuid.UUID
	NIP       string
	DateAt    time.Time
	IsIn      bool
	IsOut     bool
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (absensi *Absesnsi) BeforeCreate(tx *gorm.DB) (err error) {
	absensi.Id = uuid.New()
	return
}
