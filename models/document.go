package models

import (
	_ "fmt"
	"time"

	"github.com/google/uuid"
)

type Document struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primaryKey" example:"1"`
	AuthorID uuid.UUID `gorm:"type:uuid" json:"authorId"`
	// Author    Author    `gorm:"foreignKey:AuthorID;references:ID" json:"author"`
	Title     string    `json:"title" example:"Document Title"`
	Content   string    `json:"content" example:"Document Content"`
	CreatedAt time.Time `json:"created_at" example:"2020-09-20T14:00:00+09:00"`
	UpdatedAt time.Time `json:"updated_at" example:"2020-09-20T14:00:00+09:00"`
	CreatedBy string    `json:"created_by" example:"John"`
	UpdatedBy string    `json:"updated_by" example:"Doe"`
}
