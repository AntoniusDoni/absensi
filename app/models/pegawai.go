package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Pegawai struct {
	Id        uuid.UUID      `gorm:"primaryKey" json:"id,omitempty"`
	IdRole    uuid.UUID      `json:"id_role,omitempty"`
	Nip       string         `gorm:"index" json:"nip,omitempty"`
	Name      string         `gorm:"index" json:"name,omitempty"`
	Password  string         `json:"password,omitempty"`
	Phone     string         `json:"phone,omitempty"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:",omitempty"`
	UpdatedAt time.Time      `json:"update_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:",omitempty"`
}

func (pegawai *Pegawai) BeforeCreate(tx *gorm.DB) (err error) {
	pegawai.Id = uuid.New()
	return
}

type LoginRequest struct {
	Nip      string `json:"nip" form:"nip" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}
type LoginResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Token   string `json:"token,omitempty"`
}

type UserRole struct {
	Id       uuid.UUID `gorm:"primaryKey" json:"id,omitempty"`
	Nip      string    `gorm:"index" json:"nip,omitempty"`
	Name     string    `gorm:"index" json:"name,omitempty"`
	Password string    `json:"password,omitempty"`
	IdRole   uuid.UUID `json:"id_role,omitempty"`
	RoleName string    `json:"roles_name," `
}
