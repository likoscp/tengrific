package main

import (
	"log"
	"net/http"
	"tengrific/config"
	"tengrific/routes"
)

func main() {
	config.LoadEnv()

	router := routes.InitRoutes()
	log.Println("Server running on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
