package controller

import (
	"github.com/7lawa9111/Project-Based-Mentorship-Go-App/issues/database/connection"
	"github.com/7lawa9111/Project-Based-Mentorship-Go-App/issues/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetDocuments(c *gin.Context) {
	documents := []models.Document{}
	connection.DB.Find(&documents)
	c.String(http.StatusOK, "Hello World")
}
