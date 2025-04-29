package main

import (
	"fmt"
	"github.com/7lawa9111/Project-Based-Mentorship-Go-App/issues/models"
	"github.com/google/uuid"
)

func main() {
	fmt.Println(models.Author{
		ID:        uuid.New(),
		FirstName: "John",
		LastName:  "Doe",
	})
}
