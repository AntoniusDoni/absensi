package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Absensi struct {
	Id        uuid.UUID `gorm:"primaryKey"`
	IdPegawai uuid.UUID
	Nip       string
	DateAt    time.Time      `gorm:"type:date"`
	TimeIn    datatypes.Time `gorm:"type:time"`
	TimeOut   datatypes.Time `gorm:"type:time"`
	IsIn      bool
	IsOut     bool
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (absensi *Absensi) BeforeCreate(tx *gorm.DB) (err error) {
	absensi.Id = uuid.New()
	return
}

type RequestAbsen struct {
	Nip    string    `json:"nip" form:"nip" validate:"required"`
	Long   string    `json:"long" form:"long" validate:"required"`
	Lat    string    `json:"lat" form:"lat" validate:"required"`
	DateAt time.Time `json:"date_at" form:"lat" validate:"required"`
	IsIn   bool      `json:"in" form:"lat" validate:"required"`
	IsOut  bool      `json:"out" form:"lat" validate:"required"`
}
type ResponseRequest struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data,omitempty"`
}
