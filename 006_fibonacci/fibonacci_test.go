package fibonacci

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFibonacci(t *testing.T) {
	fibonacciTest := []struct {
		Name            string
		Position        int64
		FibonacciNumber int64
		Error           bool
	}{
		{
			Name:            "Given the position of the first number in the Fibonacci sequence should return 1",
			Position:        1,
			FibonacciNumber: 1,
			Error:           false,
		},
		{
			Name:            "Given the any number below or equal 0 should return 0",
			Position:        0,
			FibonacciNumber: 0,
			Error:           false,
		},
		{
			Name:            "Given the position of the forth number in the Fibonacci sequence should return 3",
			Position:        4,
			FibonacciNumber: 3,
			Error:           false,
		},
		{
			Name:            "Given the position 92 should be able to calculate based on int64 limitation",
			Position:        92,
			FibonacciNumber: 7540113804746344448,
			Error:           false,
		},
		{
			Name:            "Given the position 93 or more should return overflow error and return 0",
			Position:        93,
			FibonacciNumber: 0,
			Error:           true,
		},
		{
			Name:            "Given negative number should return 0 and outside of sequence error",
			Position:        -2,
			FibonacciNumber: 0,
			Error:           true,
		},
	}

	for _, tt := range fibonacciTest {
		t.Run(tt.Name, func(t *testing.T) {
			got, err := Fibonacci(tt.Position)

			assert.Equal(t, tt.Error, err != nil)
			assert.Equal(t, tt.FibonacciNumber, got)
		})
	}

}
