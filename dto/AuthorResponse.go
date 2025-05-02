package dto

import "github.com/7lawa9111/Project-Based-Mentorship-Go-App/models"

type GetAuthorsResponse struct {
	Data       []models.Author    `json:"data"`
	Pagination PaginationResponse `json:"pagination"`
}
