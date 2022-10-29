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
		User:         dbService.User,
		Institute:    dbService.Institute,
		Event:        dbService.Event,
		Reward:       dbService.Reward,
		Trainee:      dbService.Trainee,
		Team:         dbService.Team,
		Solution:     dbService.Solution,
		Entrepreneur: dbService.Entrepreneur,
	}

	app.App(base)
}
