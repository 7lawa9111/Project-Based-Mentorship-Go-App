package main

import (
	"Project-Based-Mentorship-Go-App/models"
	"fmt"
	"github.com/google/uuid"
)

func main() {
	fmt.Println(models.Author{
		id:        uuid.New(),
		firstName: "John",
		lastName:  "Doe",
	})
}
