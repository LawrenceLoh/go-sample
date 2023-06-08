package repository

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"sample/config"
)

var (
	DbInstance *gorm.DB
)

func ConnectDatabase() (*gorm.DB, error) {
	fmt.Println("Set up postgresql instance ....")
	dsn := config.SysConfig.PostgresSource
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DbInstance = db
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
		return nil, err
	}
	fmt.Println("PostgreSQL connected successfully...")
	return DbInstance, nil
}
