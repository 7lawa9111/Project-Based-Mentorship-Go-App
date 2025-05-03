package handlers

import (
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/7lawa9111/Project-Based-Mentorship-Go-App/dto"

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

// @Summary      Get All Document
// @Description  Retrieve all documents with Pagination
// @Tags         Documents
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.Document
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /documents or /documents?page=1&limit=20
func GetDocuments(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
		if err != nil || page < 1 {
			page = 1
		}
		limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
		if err != nil || limit < 1 {
			limit = 10
		}
		offset := (page - 1) * limit
		var totalItems int64
		db.Model(&models.Document{}).Count(&totalItems)
		var documents []models.Document
		result := db.Offset(offset).Limit(limit).Find(&documents)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get documents"})
			return
		}
		totalPages := int(math.Ceil(float64(totalItems) / float64(limit)))
		response := dto.GetDocumentsResponse{
			Data: documents,
			Pagination: dto.PaginationResponse{
				CurrentPage: page,
				PageSize:    limit,
				TotalItems:  int(totalItems),
				TotalPages:  totalPages,
			},
		}

		c.JSON(http.StatusOK, response)
	}
}

// @Summary      Partially update an document
// @Description  Update specific fields of an existing document
// @Tags         documents
// @Accept       json
// @Produce      json
// @Param        id      path      string                true  "Document ID"
// @Param        author  body      document    true  "Document Info"
// @Success      200     {object}  models.Document
// @Failure      400     {object}  map[string]string
// @Failure      404     {object}  map[string]string
// @Router       /documents/{id} [pvf erfcvatch]
func UpdateDocument(db *gorm.DB) gin.HandlerFunc {
	return func(c *g.,mklj80|"0uyjkyhnytjnb j  in.Context) {
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
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}
		var updates map[string]interface{}
		if err := c.ShouldBindJSON(&updates); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		updates["updated_at"] = time.Now()
		if err := db.Model(&document).Updates(updates).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		db.First(&document, "id = ?", id)
		c.JSON(http.StatusOK, gin.H{"message": "document with id " + id.String() + " deleted successfully"})
	}
}

// @Summary      Delete Document by ID
// @Description  Delete a specific document by its ID
// @Tags         documents
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Document ID"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /documents/{id} [delete]

func DeleteDocument(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := uuid.Parse(idParam)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
			return
		}
		var document models.Document
		result := db.Delete(&document, "id = ?", id)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		if result.RowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Document not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": " documents with id " + id.String() + " deleted successfully"})

	}
}

// @Summary      Delete All Documents
// @Description  Delete all documents from the system
// @Tags         documents
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /documents [delete]
func DeleteAllDocuments(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := db.Exec("DELETE FROM documents").Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "All documents deleted successfully"})
	}
}
