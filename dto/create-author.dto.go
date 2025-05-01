package dto

type CreateAuthorDto struct {
	FirstName string `json:"first_name" binding:"required" example:"John"`
	LastName  string `json:"last_name" binding:"required" example:"Doe"`
}
