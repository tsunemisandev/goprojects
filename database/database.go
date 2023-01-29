package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/tsunemisandev/goproject/models"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb(){
	dsn := fmt.Spring("host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Tokyo", 
	os.Getenv("DB_USER"),
	os.Getenv("DB_PASSWORD"),
	os.Getenv("DB_NAME"))

	db, err = gorm.Open(postgres.Open(dsn),
	&gorm.Config(Logger: logger.Default.LogMode(logger.Info))
	)

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		ost.Exit(2)
	}

	log.Println("Connected")
	db.Logger = logger.Dfault.LogMode(log.info)
	db.AutoMigrate(&models.Fact{})

	DB = Dbinstance(Db: db,)
}