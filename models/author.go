package models

import (
	_ "fmt"
	"github.com/google/uuid"
	"time"
)

type Author struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	CreatedAt time.Time
	UpdatedAt time.Time
	CreatedBy string
	UpdatedBy string
}
