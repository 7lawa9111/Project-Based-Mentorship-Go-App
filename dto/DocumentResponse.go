package dto

import "github.com/7lawa9111/Project-Based-Mentorship-Go-App/models"

type PaginationResponse struct {
	CurrentPage int `json:"current_page"`
	PageSize    int `json:"page_size"`
	TotalItems  int `json:"total_items"`
	TotalPages  int `json:"total_pages"`
}

type GetDocumentsResponse struct {
	Data       []models.Document  `json:"data"`
	Pagination PaginationResponse `json:"pagination"`
}
