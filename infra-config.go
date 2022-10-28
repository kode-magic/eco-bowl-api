package main

import "github.com/kode-magic/eco-bowl-api/infra"

func ConnectDB() *infra.Repositories {
	dbService, dbErr := infra.DBConfiguration()

	if dbErr != nil {
		panic(dbErr)
	}

	errAuto := dbService.AutoMigrate()

	if errAuto != nil {
		panic(errAuto.Error())
	}

	return dbService
}
