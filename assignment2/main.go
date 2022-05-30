package main

import (
	"assignment2/database"
	errorhandler "assignment2/errorHandler"
	"assignment2/models"
	"assignment2/route"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	content, err := ioutil.ReadFile("./config.json")
	errorhandler.ErrorChecker(err)

	var data models.Config
	err = json.Unmarshal(content, &data)
	errorhandler.ErrorChecker(err)

	fmt.Println(data)

	database.StartDB(data.DbHost, data.DbPort, data.DbUser, data.DbPass, data.DbName)

	route.StartServer().Run(data.ServicePort)
}
