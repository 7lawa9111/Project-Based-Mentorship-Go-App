package routes

import (
	"github.com/7lawa9111/Project-Based-Mentorship-Go-App/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterDocumentsRoutes(router *gin.Engine, db *gorm.DB) {
	documentGroup := router.Group("/documents")
	{
		documentGroup.POST("", handlers.CreateDocument(db))
		documentGroup.GET("/:id", handlers.GetDocumentByID(db))
		documentGroup.GET("", handlers.GetDocuments(db))
		documentGroup.PATCH("/:id", handlers.UpdateDocument(db))
		documentGroup.DELETE("/:id", handlers.DeleteDocument(db))
		documentGroup.DELETE("", handlers.DeleteAllDocuments(db))
	}
}
