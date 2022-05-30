package service

import (
	"assignment2/database"
	errorhandler "assignment2/errorHandler"
	"assignment2/models"
	"fmt"
)

func CreateOrder(data *models.Request) bool {
	err, id := database.CreateOrder(&data.Order)
	errorhandler.ErrorChecker(err)
	for i := 0; i < len(data.Items); i++ {
		data.Items[i].OrderID = id
	}
	err = database.CreateItem(&data.Items)
	errorhandler.ErrorChecker(err)

	return true
}

func GetOrders() []models.Request {
	orders := database.GetOrders()
	var res []models.Request

	for i := 0; i < len(orders); i++ {
		var items []models.Item = database.GetItems(orders[i].ID)
		res = append(res, models.Request{
			Order: orders[i],
			Items: items,
		})
	}

	return res
}

func UpdateOrder(orderId uint, req *models.Request) bool {
	order := req.Order
	order.ID = orderId
	items := req.Items
	fmt.Println(order)

	isSucces := database.UpdateOrder(&order)

	for i := 0; i < len(items); i++ {
		items[i].OrderID = orderId
	}
	err := database.CreateItem(&items)
	errorhandler.ErrorChecker(err)

	return isSucces
}

func DeleteOrder(id uint) bool {
	isSucces := database.DeleteOrder(id)

	return isSucces
}
