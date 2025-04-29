package models

import (
	_ "fmt"
	"github.com/google/uuid"
	"time"
)

type Document struct {
	ID        uuid.UUID
	AuthorId  uuid.UUID
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	CreatedBy string
	UpdatedBy string
}
