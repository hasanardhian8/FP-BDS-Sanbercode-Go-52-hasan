package main

import (
	"member/config"
	"member/routes"
)

func main() {
	db := config.GetDatabaseConnection()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	r := routes.Router(db)
	r.Run()
}
