package calculator

import "errors"

func Addition(a float64, b float64) float64 {
	return a + b
}

func Subtraction(a float64, b float64) float64 {
	return a - b
}

func Division(a float64, b float64) (float64, error) {
	if b == 0 || a == 0 {
		return 0, errors.New("cannot divide by or a zero")
	}
	return a / b, nil
}

func Multiplication(a float64, b float64) float64 {
	return a * b
}
