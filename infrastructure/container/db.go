package container

import (
	"go-microservice/domain/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func NewGormDBConnection(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error opening database with GORM: %v", err)
	}

	return db
}

func InitializeTables(db *gorm.DB) {
	err := db.AutoMigrate(&entities.User{})
	if err != nil {
		return
	}
}
