package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Cuti struct {
	Id        uuid.UUID `gorm:"primaryKey"`
	IdPegawai uuid.UUID
	DateAt    time.Time
	DateEnd   time.Time
	IdAtasan  uuid.UUID
	IsApprove bool
}

func (cuti *Cuti) BeforeCreate(tx *gorm.DB) (err error) {
	cuti.Id = uuid.New()
	return
}
