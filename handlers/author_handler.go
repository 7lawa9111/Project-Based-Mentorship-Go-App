package handlers

import (
	"github.com/7lawa9111/Project-Based-Mentorship-Go-App/dto"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/7lawa9111/Project-Based-Mentorship-Go-App/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// @Summary      Create a new author
// @Description  Add a new author to the system
// @Tags         authors
// @Accept       json
// @Produce      json
// @Param        author  body      dto.CreateAuthorDto  true  "Author Info"
// @Success      201     {object}  models.Author
// @Failure      400     {object}  map[string]string
// @Router       /authors [post]
func CreateAuthor(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.Author

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

// @Summary      Get author by ID
// @Description  Retrieve an author's information by their unique ID
// @Tags         authors
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Author ID"
// @Success      200  {object}  models.Author
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /authors/{id} [get]
func GetAuthorByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := uuid.Parse(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
			return
		}

		var author models.Author
		if err := db.First(&author, "id = ?", id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, author)
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
func GetAuthors(db *gorm.DB) gin.HandlerFunc {
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
		db.Model(&models.Author{}).Count(&totalItems)
		var authors []models.Author
		result := db.Offset(offset).Limit(limit).Find(&authors)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get authors"})
			return
		}
		totalPages := int(math.Ceil(float64(totalItems) / float64(limit)))
		response := dto.GetAuthorsResponse{
			Data: authors,
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

// @Summary      Partially update an author
// @Description  Update specific fields of an existing author
// @Tags         authors
// @Accept       json
// @Produce      json
// @Param        id      path      string                true  "Author ID"
// @Param        author  body      dto.PatchAuthorDto    true  "Author Info"
// @Success      200     {object}  models.Author
// @Failure      400     {object}  map[string]string
// @Failure      404     {object}  map[string]string
// @Router       /authors/{id} [patch]
func UpdateAuthor(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := uuid.Parse(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
			return
		}
		var author models.Author
		if err := db.First(&author, "id = ?", id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
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
		if err := db.Model(&author).Updates(updates).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		db.First(&author, "id = ?", id)
		c.JSON(http.StatusOK, author)
	}
}
