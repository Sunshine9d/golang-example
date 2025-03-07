package main

import (
	"fmt"
)

type APIError struct {
	// your code goes here
}

func (e *APIError) Error() string {
	// your code goes here
}

func NewAPIError(code int, message string, details map[string]interface{}) error {
	// your code goes here
}

func main() {
	err := NewAPIError(400, "Bad request", map[string]interface{}{
		"field": "username",
		"error": "cannot be empty",
	})
	fmt.Println(err.Error())
}
