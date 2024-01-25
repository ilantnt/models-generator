package main

import (
	"fmt"
	"models_generator/generated-models"
)

func main() {
	// Example using the generated models
	person := generated_models.Person{
		Name:      "John Doe",
		Age:       30,
		Addresses: []generated_models.Address{{City: "New York", Country: "USA"}},
	}
	fmt.Println(person)

	// Example Student
	student := generated_models.Student{
		Classes: []string{"Math", "Science"},
		Grade:   95.5,
	}
	fmt.Println(student)
}
