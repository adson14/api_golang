package main

import (
	"api_golang/database"
	"api_golang/server"
)

func main() {

	database.StartDb()
	server := server.NewServer()

	server.Run()
}
