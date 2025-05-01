package main

import (
	"github.com/7lawa9111/Project-Based-Mentorship-Go-App/issues/database/connection"
	"github.com/7lawa9111/Project-Based-Mentorship-Go-App/issues/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	er := godotenv.Load()
	if er != nil {
		log.Fatal(" Error loading .env file")
	}
	err := connection.NewDatabaseConnection()
	if err != nil {
		log.Fatal(" Failed to connect to database:", err)
	}

	router := gin.New()
	routes.DocumentRoute(router)
	router.Run(":8080")

}
