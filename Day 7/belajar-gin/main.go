package main

import (
	"belajar-gin/router"
)

func main() {
	var PORT = ":8080"

	router.StartServer().Run(PORT)
}
