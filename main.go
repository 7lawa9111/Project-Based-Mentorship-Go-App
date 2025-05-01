package main

import (
	"log"

	"database/connection"
)

func main() {
	db, err := connection.NewDatabaseConnection()
	if err != nil {
		log.Fatal(" Failed to connect to database:", err)
	}

	log.Println(" Connected to database and migrations done!")

}
