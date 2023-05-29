package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var lock = &sync.Mutex{}

type GormDB struct {
	Db   *gorm.DB
	Conn *sql.DB
}

func New() *GormDB {
	return &GormDB{}
}

func (godb *GormDB) GetInstanceConnect() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println("Driver", os.Getenv("DB_Driver"))
	if os.Getenv("DB_Driver") == "mysql" {
		// destination:=fmt.Sprintf("user")
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			os.Getenv("DB_User"),
			os.Getenv("DB_Password"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_Port"),
			os.Getenv("DB_Name"),
		)
		lock.Lock()
		var err error
		godb.Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		lock.Unlock()
		if err != nil {
			// log.Fatal(err)
			log.Printf("Error %s when connection DB\n", err)

			godb.CreateDB()
		}
	} else {
		destination := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_Port"),
			os.Getenv("DB_User"),
			os.Getenv("DB_Password"),
			os.Getenv("DB_Name"),
			os.Getenv("DB_TimeZone"),
		)

		lock.Lock()
		godb.Db, err = gorm.Open(postgres.Open(destination), &gorm.Config{})
		lock.Unlock()
		if err != nil {
			//
			log.Printf("Error %s when connection DB\n", err)
			log.Fatal(err)

		}
	}

	log.Printf("DB Connect")
	return godb.Db, nil
}
func (godb *GormDB) CreateDB() {
	godotenv.Load()
	if os.Getenv("DB_Driver") == "mysql" {
		destcon := fmt.Sprintf("%s:%s@tcp(%s:%s)/", os.Getenv("DB_User"), os.Getenv("DB_Password"), os.Getenv("DB_HOST"), os.Getenv("DB_Port"))
		db, err := sql.Open("mysql", destcon)
		if err != nil {
			log.Printf("Error %s when connection DB\n", err)
			log.Fatal(err)
		}
		dbname := fmt.Sprintf("CREATE DATABASE %s", os.Getenv("DB_Name"))
		_, err = db.Exec(dbname)
		if err != nil {
			log.Printf("Error %s when Create DataBase\n", err)
			log.Fatal(err)
		}
		log.Printf("DB Create")
		godb.GetInstanceConnect()
	}
}
