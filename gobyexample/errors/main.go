package main

import (
	"errors"
	"fmt"
)



func validateInput(input string) error {
	if input == "" {
		return fmt.Errorf("%w: input is empty", ErrValidation)
	}
	return nil
}
func findRecord(id int) error {
	// Simulate database lookup
	if id != 42 {
		return fmt.Errorf("%w: record with id %d", ErrNotFound, id)
	}
	return nil
}
func processData(input string, id int) error {
	err := validateInput(input)
	if err != nil {
		return fmt.Errorf("input validation failed: %w", err)
	}
	err = findRecord(id)
	if err != nil {
		return fmt.Errorf("finding record failed: %w", err)
	}
	return nil
}
func main() {
	err := processData("", 123) // Empty input
	if err != nil {
		fmt.Println("Final error:", err)
		// Unwrap the error chain
		for unwrappedErr := err; unwrappedErr != nil; unwrappedErr = errors.Unwrap(unwrappedErr) {
			fmt.Println("  Caused by:", unwrappedErr)
		}
		if errors.Is(err, ErrValidation) {
			fmt.Println("it is a validation error")
		}
	}
	err2 := processData("ok", 42)
	if err2 == nil {
		fmt.Println("Everything is okay!")
	}
	err3 := processData("ok", 1)
	if err3 != nil {
		if errors.Is(err3, ErrNotFound) {
			fmt.Println("Record has not been found.")
		}
	}
}
