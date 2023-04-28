package main

import (
	"testing"
)

func TestPerformOperation(t *testing.T) {
	t.Run("addition", func(t *testing.T) {
		result, err := performOperation(5.0, 3.0, "+")
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if result != 8.0 {
			t.Errorf("Expected 8.0, got %v", result)
		}
	})

	// Add more test cases for subtraction, multiplication, division, division by zero, and invalid operator
}
