package calculator

import (
	"testing"
)

func TestAddition(t *testing.T) {
	result := Addition(2, 3)
	expected := 5.0
	if result != expected {
		t.Errorf("Addition failed. Got: %v, Expected: %v", result, expected)
	}
}

func TestSubtraction(t *testing.T) {
	result := Subtraction(5, 3)
	expected := 2.0
	if result != expected {
		t.Errorf("Subtraction failed. Got: %v, Expected: %v", result, expected)
	}
}

func TestMultiplication(t *testing.T) {
	result := Multiplication(2, 3)
	expected := 6.0
	if result != expected {
		t.Errorf("Multiplication failed. Got: %v, Expected: %v", result, expected)
	}
}

func TestDivision(t *testing.T) {
	result, err := Division(6, 3)
	expected := 2.0
	if err != nil {
		t.Errorf("Division failed with error: %v", err)
	}
	if result != expected {
		t.Errorf("Division failed. Got: %v, Expected: %v", result, expected)
	}
}

func TestDivisionByZero(t *testing.T) {
	_, err := Division(6, 0)
	if err == nil {
		t.Error("Expected error for division by zero, but got none.")
	}
}
