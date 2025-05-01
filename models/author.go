package models

import (
	_ "fmt"
	"time"

	"github.com/google/uuid"
)

type Author struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey" example:"123e4567-e89b-12d3-a456-426614174000"`
	FirstName string    `json:"first_name" example:"John"`
	LastName  string    `json:"last_name" example:"Doe"`
	CreatedAt time.Time `json:"created_at" example:"2020-09-20T14:00:00+09:00"`
	UpdatedAt time.Time `json:"updated_at" example:"2020-09-20T14:00:00+09:00"`
	CreatedBy string    `json:"created_by" example:"John"`
	UpdatedBy string    `json:"updated_by" example:"Doe"`
}
