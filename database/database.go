package database

import (
	"fmt"
	"github.ademola.com/trivia-app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() {
	// dsn is Data Source Network - The Database you will be connection to
	dsn := fmt.Sprintf("host=db user=%s password%s dbname=%s port=5433 sslmode=disable TimeZone=EUROPE/Berlin",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	log.Println("Connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running migrations")
	db.AutoMigrate(&models.Fact{})

	DB = Dbinstance{
		Db: db,
	}

}
