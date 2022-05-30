package database

import (
	errorhandler "assignment2/errorHandler"
	"assignment2/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func StartDB(host, port, user, password, dbName string) {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, port)
	fmt.Println(config)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	errorhandler.ErrorChecker(err)

	db.AutoMigrate(models.Order{}, models.Item{})
}

func GetDB() *gorm.DB {
	return db
}
