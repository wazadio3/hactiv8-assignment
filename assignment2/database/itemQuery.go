package database

import (
	errorhandler "assignment2/errorHandler"
	"assignment2/models"
)

func CreateItem(item *[]models.Item) error {
	err := db.Create(&item).Error
	return err
}

func GetItems(id uint) []models.Item {
	var items []models.Item
	err := db.Where("order_id = ?", id).Find(&items).Error
	errorhandler.ErrorChecker(err)
	return items
}
