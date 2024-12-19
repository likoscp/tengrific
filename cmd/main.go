package main

import (
	"log"
	"net/http"
	"tengrific/config"
	"tengrific/routes"
	"tengrific/db"
)

func main() {
	config.LoadEnv()
	db.ConnectDB()
	router := routes.InitRoutes()
	log.Println("Server running on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
