package config

import (
	"fmt"
	"log"
	"member/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

func GetDatabaseConnection() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("POSTGRES_URL")))
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Sukses Konek ke Db!")
	db.AutoMigrate(
		&models.Produks{},
		&models.Saldos{},
		&models.Registers{},
		&models.Pemesanans{},
		&models.Profils{},
		&models.Transaksis{},
		&models.Feedbacks{})
	return db
}
