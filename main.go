package main

import (
	"github.com/kode-magic/eco-bowl-api/app"
	service "github.com/kode-magic/eco-bowl-api/services"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("app.env")

	if err != nil {
		log.Fatalf("Error loading app.env file")
	}

	dbService := ConnectDB()

	base := service.BaseService{
		User:      dbService.User,
		Institute: dbService.Institute,
		Event:     dbService.Event,
	}

	app.App(base)
}