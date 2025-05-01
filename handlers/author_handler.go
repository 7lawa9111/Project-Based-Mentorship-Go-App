package handlers

import (
	"net/http"
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
