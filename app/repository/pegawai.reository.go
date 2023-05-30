package repository

import (
	"github.com/absensi/app/models"
)

func (repo Repository) GetUserNip(nip string, user *models.UserRole) error {
	db, _ := repo.Gormdb.GetInstanceConnect()
	err := db.Model(&models.Pegawai{}).Select("pegawais.id as id", "nip", "id_role", "pegawais.name as name", "password", "phone", "roles.name as roles_name").Joins("inner join roles", "roles.id=id_role").Where("nip=? ", nip).Scan(&user).Error
	close, _ := db.DB()
	close.Close()
	return err
}
func (repo Repository) CreateUser(user *models.Pegawai) error {
	db, _ := repo.Gormdb.GetInstanceConnect()
	err := db.Create(&user).Error
	close, _ := db.DB()
	close.Close()
	return err
}
