package database

import (
	"fmt"
	"log"
	"tugas8/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



var (
	host	 = "localhost"
	port	 = 5432
	user	 = "postgres"
	password = "1234"
	dbname	 = "postgres"
	db *gorm.DB
	err error
)


func StartDB(){
	config := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database", err)

	}

	log.Println("database connected successfully")


	err = db.AutoMigrate(models.Book{})
	if err != nil {
		log.Fatal("error migrating database", err)
	}
	
	log.Print("database migrated successfully")
}

func GetDB() *gorm.DB {
	return db
}