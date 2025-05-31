package main

import (
	"hacka/api"
	"hacka/infra/database"
	"os"
)

func main() {
	db := database.StartDB()
	defer db.Close()

	myApp := api.NewApp(db)

	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}

	api.RunServer(port, myApp)
}
