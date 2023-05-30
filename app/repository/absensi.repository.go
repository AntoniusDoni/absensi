package repository

import (
	"github.com/absensi/app/models"
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

func (repo Repository) CreateAbsen(absen *models.Absensi) error {
	db, _ := repo.Gormdb.GetInstanceConnect()
	err := db.Create(&absen).Error
	close, _ := db.DB()
	close.Close()
	return err
}

func (repo Repository) GetDetailAbsen(id_pegawai uuid.UUID, dateat string, isout string) (*models.Absensi, error) {
	db, _ := repo.Gormdb.GetInstanceConnect()
	absen := new(models.Absensi)
	queryabsen := db.Model(&models.Absensi{}).Where("id_pegawai=?", id_pegawai).Where("date_at=?", dateat)
	var err error
	if isout != "" {
		queryabsen = queryabsen.Where("is_out=?", isout)
	}
	err = queryabsen.First(absen).Error
	close, _ := db.DB()
	close.Close()
	return absen, err
}

func (repo Repository) UpdateAbsen(id uuid.UUID, timeout datatypes.Time) int64 {
	db, _ := repo.Gormdb.GetInstanceConnect()
	err := db.Model(&models.Absensi{}).Where("id=?", id).Where("is_out!=1").Updates(map[string]interface{}{"time_out": timeout, "is_out": true}).RowsAffected
	close, _ := db.DB()
	close.Close()
	return err
}
