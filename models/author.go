package models

import (
	_ "fmt"
	"github.com/google/uuid"
	"time"
)

type Author struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey" example:"1"`
	FirstName string    `json:"first_name" example:"John"`
	LastName  string    `json:"last_name" example:"Doe"`
	CreatedAt time.Time `json:"created_at" example:"2020-09-20T14:00:00+09:00"`
	UpdatedAt time.Time `json:"updated_at" example:"2020-09-20T14:00:00+09:00"`
	CreatedBy string    `json:"created_by" example:"John"`
	UpdatedBy string    `json:"updated_by" example:"Doe"`
}
