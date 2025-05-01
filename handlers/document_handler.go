package handlers

import (
	"net/http"
	"time"

	"github.com/7lawa9111/Project-Based-Mentorship-Go-App/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// @Summary      Create a new Document
// @Description  Add a new Document to the system
// @Tags         Documents
// @Accept       json
// @Produce      json
// @Param        Document  body       dto.CreateDocumentDto  true  "Document Info"
// @Success      201     {object}  models.Document
// @Failure      400     {object}  map[string]string
// @Router       /documents [post]
func CreateDocument(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.Document

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		input.ID = uuid.New()
		input.CreatedAt = time.Now()
		input.UpdatedAt = time.Now()

		if err := db.Create(&input).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, input)
	}
}

// @Summary      Get Document by ID
// @Description  Retrieve a document by its unique ID
// @Tags         Documents
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Document ID"
// @Success      200  {object}  models.Document
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /documents/{id} [get]
func GetDocumentByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := uuid.Parse(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
			return
		}

		var document models.Document
		if err := db.First(&document, "id = ?", id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Document not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, document)
	}
}
