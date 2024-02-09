package main

import (
	"assigment/config"
	"assigment/controllers"
	"assigment/routers"
)

var PORT = ":8080"

func main() {
	db := config.StartDB()
	controller := controllers.New(db)

	routers.StartServer(controller).Run(PORT)
}
