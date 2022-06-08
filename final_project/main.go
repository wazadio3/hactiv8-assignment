package main

import (
	"final_project/database"
	"final_project/errorhandler"
	"final_project/router"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	errorhandler.Check(err)
	db := database.StartDB(os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	router.StartService(db).Run(os.Getenv("SERVICE_PORT"))
}
