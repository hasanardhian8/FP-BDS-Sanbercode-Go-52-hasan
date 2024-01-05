package main

import (
	"member/config"
	"member/routes"
)

func Main() {
	db := config.GetDatabaseConnection()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	r := routes.Router(db)
	r.Run()
}
