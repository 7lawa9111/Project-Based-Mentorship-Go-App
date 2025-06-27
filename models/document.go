package models

import (
	_ "fmt"
	"time"

	"github.com/google/uuid"
)

type Document struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primaryKey" example:"123e4567-e89b-12d3-a456-426614174000"`
	AuthorID uuid.UUID `gorm:"type:uuid" json:"authorId"`
	// Author    Author    `gorm:"foreignKey:AuthorID;references:ID" json:"author"`
	Title     string    `json:"title" example:"Document Title"`
	Content   []byte    `json:"content" example:"RG9jdW1lbnQgQ29udGVudA=="` // Changed from string to []byte
	CreatedAt time.Time `json:"created_at" example:"2020-09-20T14:00:00+09:00"`
	UpdatedAt time.Time `json:"updated_at" example:"2020-09-20T14:00:00+09:00"`
}
