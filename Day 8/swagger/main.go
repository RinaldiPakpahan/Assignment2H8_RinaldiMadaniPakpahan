package main

import (
	"swagger/database"
	"swagger/router"
)

func main() {
	var PORT = ":8080"

	database.StartDB()

	router.StartServer().Run(PORT)
}
