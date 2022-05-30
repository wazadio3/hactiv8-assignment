package controller

import (
	errorhandler "assignment2/errorHandler"
	"assignment2/models"
	"assignment2/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var payload models.Request

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	result := service.CreateOrder(&payload)

	if result {
		c.JSON(http.StatusCreated, gin.H{
			"Message": "OK",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "NOT OK",
		})
	}
}

func GetOrders(c *gin.Context) {
	res := service.GetOrders()

	c.JSON(http.StatusAccepted, res)
}

func UpdateOrder(c *gin.Context) {
	var payload models.Request

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	orderId, err := strconv.ParseUint(c.Param("orderId"), 10, 32)
	errorhandler.ErrorChecker(err)
	id := uint(orderId)

	result := service.UpdateOrder(id, &payload)

	if result {
		c.JSON(http.StatusCreated, gin.H{
			"Message": "OK",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "NOT OK",
		})
	}
}

func DeleteOrder(c *gin.Context) {
	orderId, err := strconv.ParseUint(c.Param("orderId"), 10, 32)
	errorhandler.ErrorChecker(err)
	id := uint(orderId)

	result := service.DeleteOrder(id)

	if result {
		c.JSON(http.StatusCreated, gin.H{
			"Message": "OK",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "NOT OK",
		})
	}
}
