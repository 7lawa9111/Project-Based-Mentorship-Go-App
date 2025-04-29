package models

import (
	_ "fmt"
	"github.com/google/uuid"
	"time"
)

type Author struct {
	id        uuid.UUID
	firstName string
	lastName  string
	createdAt time.Time
	updatedAt time.Time
	createdBy string
	updatedBy string
}
