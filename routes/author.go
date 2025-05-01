package routes

import (
	"github.com/7lawa9111/Project-Based-Mentorship-Go-App/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAuthorRoutes(router *gin.Engine, db *gorm.DB) {
	authorGroup := router.Group("/authors")
	{
		authorGroup.POST("", handlers.CreateAuthor(db))
	}
}
