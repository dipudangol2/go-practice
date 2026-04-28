package main

import (
	"errors"
	"fmt"
)

// Sentinel error
var ErrDivisionByZero = errors.New("division by zero")

// Custom error type
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error on %q: %s", e.Field, e.Message)
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}
	return a / b, nil
}

func validateAge(age int) error {
	if age < 0 {
		return &ValidationError{Field: "age", Message: "must be non-negative"}
	}
	if age > 150 {
		return &ValidationError{Field: "age", Message: "unrealistically large"}
	}
	return nil
}

// Wrapping errors (Go 1.13+)
func processAge(age int) error {
	if err := validateAge(age); err != nil {
		return fmt.Errorf("processAge(%d): %w", age, err)
	}
	return nil
}

func main() {
	// Basic error handling
	result, err := divide(10, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 3 = %.4f\n", result)
	}

	// Sentinel error check
	_, err = divide(5, 0)
	if errors.Is(err, ErrDivisionByZero) {
		fmt.Println("Caught sentinel error:", err)
	}

	// Custom error type
	err = validateAge(-1)
	var valErr *ValidationError
	if errors.As(err, &valErr) {
		fmt.Printf("Field: %s, Problem: %s\n", valErr.Field, valErr.Message)
	}

	// Wrapped errors
	err = processAge(200)
	if err != nil {
		fmt.Println("Wrapped error:", err)
		// Unwrap to find the original ValidationError
		var ve *ValidationError
		if errors.As(err, &ve) {
			fmt.Println("Underlying field:", ve.Field)
		}
	}

	// defer / recover for panics
	fmt.Println("Result of safeDivide(8, 0):", safeDivide(8, 0))
	fmt.Println("Result of safeDivide(8, 2):", safeDivide(8, 2))
}

func safeDivide(a, b int) (result int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
			result = -1
		}
	}()
	result = a / b
	return
}
