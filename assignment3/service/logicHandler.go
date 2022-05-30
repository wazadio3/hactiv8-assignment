package service

import (
	errorhandler "assignment3/error_handler"
	"assignment3/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetData(c *gin.Context) {
	content, err := ioutil.ReadFile("./data.json")
	errorhandler.ErrorChecker(err)

	data := models.Data{}

	err = json.Unmarshal(content, &data)
	fmt.Println(data)

	res := models.Res{}
	res.Water = data.Status.Water + "m"
	res.Wind = data.Status.Wind + "m/s"
	water, err := strconv.Atoi(data.Status.Water)
	errorhandler.ErrorChecker(err)
	wind, err := strconv.Atoi(data.Status.Wind)

	if water < 5 || wind < 6 {
		res.Status = "aman"
	} else if water <= 8 || wind <= 15 {
		res.Status = "siaga"
	} else {
		res.Status = "bahaya"
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"water":  res.Water,
		"wind":   res.Wind,
		"status": res.Status,
	})
}
