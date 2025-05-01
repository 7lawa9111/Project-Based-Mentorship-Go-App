package dto

import (
	"github.com/google/uuid"
)

type CreateDocumentDto struct {
	AuthorID uuid.UUID `json:"authorId" binding:"required" example:"123e4567-e89b-12d3-a456-426614174000"`
	Title    string    `json:"title" binding:"required" example:"Document Title"`
	Content  string    `json:"content" binding:"required"  example:"Document Content"`
}
