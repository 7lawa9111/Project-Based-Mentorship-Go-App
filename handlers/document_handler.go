package handlers

import (
	"io"
	
	"math"
net/http"
	
	"net/http"
	"io"
	"path/filepath"

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
		var input dto.CreateDocumentDto

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		document := models.Document{
			ID:        uuid.New(),
			AuthorID:  input.AuthorID,
			Title:     input.Title,
			Content:   input.Content,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		if err := db.Create(&document).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, document)
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

// @Summary      Update a document
// @Description  Update an existing document
// @Tags         documents
// @Accept       json
// @Produce      json
// @Param        id          path      string                true  "Document ID"
// @Param        document    body      dto.CreateDocumentDto    true  "Document Info"
// @Success      200     {object}  models.Document
// @Failure      400     {object}  map[string]string
// @Failure      404     {object}  map[string]string
// @Router       /documents/{id} [patch]
func UpdateDocument(db *gorm.DB) gin.HandlerFunc {
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
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}

		}
		
		var input dto.CreateDocumentDto
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		}
		
		// Update document fields
		document.Title = input.Title
		document.AuthorID = input.AuthorID

		document.UpdatedAt = time.Now()
		
		if err := db.Save(&document).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		}
		
		c.JSON(http.StatusOK, document)
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

		c.JSON(http.StatusOK, gin.H{"message": "Document with id " + id.String() + " deleted successfully"})
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

// @Summary      Upload PDF document
// @Description  Upload a PDF document using multipart/form-data
// @Tags         documents
// @Accept       multipart/form-data
// @Produce      json
// @Param        authorId   formData  string  true  "Author ID"
// @Param        title      formData  string  true  "Document Title"
// @Param        file       formData  file    true  "PDF file to upload"
// @Success      201  {object}  models.Document
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /documents/upload [post]
func UploadDocument(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse form data
		authorIDStr := c.PostForm("authorId")
		title := c.PostForm("title")

		// Validate author ID
		authorID, err := uuid.Parse(authorIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid author ID format"})
			return
		}

		// Get the file from form data
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "File upload failed: " + err.Error()})
			return
		}

		// Check file extension (only allow PDF)
		ext := filepath.Ext(file.Filename)
		if ext != ".pdf" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Only PDF files are allowed"})
			return
		}

		// Open the uploaded file
		openedFile, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open uploaded file"})
			return
		}
		defer openedFile.Close()

		// Read file content
		fileContent, err := io.ReadAll(openedFile)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file content"})
			return
		}

		// Create document
		document := models.Document{
			ID:        uuid.New(),
			AuthorID:  authorID,
			Title:     title,
			Content:   fileContent,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		// Save to database
		if err := db.Create(&document).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save document: " + err.Error()})
			return
		}

		// Return success response
		c.JSON(http.StatusCreated, document)
	}
}

// @Summary      Download PDF document
// @Description  Download a PDF document by ID
// @Tags         documents
// @Produce      application/pdf
// @Param        id   path      string  true  "Document ID"
// @Success      200  {file}    binary
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /documents/{id}/download [get]
func DownloadDocument(db *gorm.DB) gin.HandlerFunc {
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

		// Set response headers for file download
		fileName := document.Title + ".pdf"
		c.Header("Content-Disposition", "attachment; filename="+fileName)
		c.Header("Content-Type", "application/pdf")
		c.Header("Content-Length", strconv.Itoa(len(document.Content)))

		// Write the PDF content to the response
		c.Writer.Write(document.Content)
	}
}
