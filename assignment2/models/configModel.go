package models

type Config struct {
	ServicePort string `json:"SERVICE_PORT"`
	DbHost      string `json:"DB_HOST"`
	DbPort      string `json:"DB_PORT"`
	DbUser      string `json:"DB_USER"`
	DbPass      string `json:"DB_PASS"`
	DbName      string `json:"DB_NAME"`
}
