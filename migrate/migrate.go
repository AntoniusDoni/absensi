package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/absensi/app/models"
	"github.com/absensi/app/services"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

func main() {
	services := services.New()
	services.Db.GetInstanceConnect()
	db, _ := services.Db.GetInstanceConnect()
	db.AutoMigrate(
		&models.Pegawai{},
		&models.Roles{},
		&models.Absensi{},
		&models.Jadwal{},
		&models.Cuti{},
	)
	Seed(db)
	con, _ := db.DB()
	con.Close()

}

func Seed(db *gorm.DB) {

	role := new(models.Roles)
	pegawai := new(models.Pegawai)
	role.Name = "SuperAdmin"
	db.Where(models.Roles{Name: role.Name}).Attrs(models.Roles{Name: role.Name}).FirstOrCreate(&role)
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	db.Where(models.Pegawai{Nip: "superadmin"}).Attrs(models.Pegawai{Nip: "superadmin", Name: "Super Admin", Password: string(hashedPassword), IdRole: role.Id}).FirstOrCreate(&pegawai)
	getDays(db)
	con, _ := db.DB()
	con.Close()
}

func getDays(db *gorm.DB) {
	days := daysInMonth(time.Now().Local())

	jadwal := new(models.Jadwal)
	for _, day := range days {
		if day != "" {
			date, err := time.Parse("2006-01-02", day)
			if err != nil {
				log.Printf("Error %s when Seed jadwal DB\n", fmt.Sprintf("%v-%v", err, day))
			}
			startTime := datatypes.NewTime(9, 0, 0, 0)
			endTime := datatypes.NewTime(17, 0, 0, 0)
			db.Where(models.Jadwal{SchaduleDate: date}).Attrs(models.Jadwal{SchaduleDate: date, SchaduleTimeIn: startTime, SchaduleTimeOut: endTime}).FirstOrCreate(&jadwal)
		}

	}

}

func daysInMonth(t time.Time) []string {
	t = time.Date(t.Year(), t.Month(), 32, 0, 0, 0, 0, time.UTC)
	daysInMonth := 32 - t.Day()
	days := make([]string, daysInMonth)
	mouth := int(t.Month())
	for i := range days {
		day := i + 1
		curdate := strconv.Itoa(day)
		curmoth := strconv.Itoa(mouth)
		if day < 10 {
			curdate = fmt.Sprintf("0%v", day)
		}
		if mouth < 10 {
			curmoth = fmt.Sprintf("0%v", mouth)
		}
		date := fmt.Sprintf("%s-%s-%s", strconv.Itoa(t.Year()), curmoth, curdate)
		if date != "2023-06-31" {
			t, err := time.Parse("2006-01-02", date)
			if err != nil {
				panic(err)
			}
			if t.Weekday() != time.Saturday && t.Weekday() != time.Sunday {
				days[i] = date
			}
		}

	}

	return days
}
