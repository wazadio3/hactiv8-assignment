package main

import (
	errorhandler "assignment3/error_handler"
	"assignment3/models"
	"assignment3/route"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

func main() {
	go updateData()
	route.StartServer().Run(":9000")
}

func updateData() {
	for {
		time.Sleep(time.Second * 5)
		newData := models.Data{}
		water := rand.Intn(101)
		wind := rand.Intn(101)

		newData.Status.Water = fmt.Sprint(water)
		newData.Status.Wind = fmt.Sprint(wind)

		content, err := json.Marshal(newData)
		errorhandler.ErrorChecker(err)
		err = ioutil.WriteFile("./data.json", content, 0644)
		errorhandler.ErrorChecker(err)
	}
}
