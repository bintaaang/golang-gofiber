package database

import (
	"fmt"
	//"gofiber-endpoint/encripyt"
	"gofiber-endpoint/encripyt"
	"log"
	"os"
	"time"
	

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	UsingPostgre       *gorm.DB
)

func InitAllDBs() {
	UsingPostgre = initSingleDB()
}

func initSingleDB() *gorm.DB {
	user := os.Getenv("USER")
	pass := os.Getenv("PASSWORD")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	name := os.Getenv("DBNAME")

	result, err := encripyt.DecryptFromBase64(name, "xQXIBUi/bZU1ZbFeFFbooHV6QxE8okplr5kxKMTOzR0="); if err != nil {
		log.Fatalf("ini adalah pesan erromu: %v", err)
	}
	fmt.Println(result)
	
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		host, user, pass, result, port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get sql.DB for: %v", err)
	}

	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(10 * time.Minute)

	return db
}
