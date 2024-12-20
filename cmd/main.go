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
	log.Println("Server running on http://127.0.0.1:8080")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", router))
}
