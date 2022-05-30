package database

import (
	errorhandler "assignment2/errorHandler"
	"assignment2/models"
)

func CreateOrder(order *models.Order) (error, uint) {
	err := db.Create(order)
	return err.Error, order.ID
}

func GetOrders() []models.Order {
	var orders []models.Order
	err = db.Find(&orders).Error
	errorhandler.ErrorChecker(err)

	return orders
}

func UpdateOrder(order *models.Order) bool {
	err := db.Save(order).Error
	errorhandler.ErrorChecker(err)
	return true
}

func DeleteOrder(id uint) bool {
	var order models.Order
	err := db.Delete(&order, id).Error
	errorhandler.ErrorChecker(err)

	return true
}
