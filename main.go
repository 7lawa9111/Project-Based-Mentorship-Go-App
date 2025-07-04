package main

import (
	"log"

	"github.com/7lawa9111/Project-Based-Mentorship-Go-App/database/connection"
	"github.com/7lawa9111/Project-Based-Mentorship-Go-App/routes"
	"github.com/gin-gonic/gin"

	// Swagger docs
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Go Project Mentorship API
// @version         1.0
// @description     This is the backend API for the go project-based mentorship platform.
// @host            localhost:8080
// @BasePath        /
func main() {
	db, err := connection.NewDatabaseConnection()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	log.Println("Connected to database and migrations done!")
	_ = db
	r := gin.Default()

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Register all route groups
	routes.RegisterAuthorRoutes(r, db)
	routes.RegisterDocumentsRoutes(r, db)

	r.Run("0.0.0.0:8080") // listen and serve

}
