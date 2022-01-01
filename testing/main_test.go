package main

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {

	tests := [][]int{
		{1, 2, 3},
		{3, 5, 8},
	}

	for _, test := range tests {

		tag := fmt.Sprintf("table testing %d+%d", test[0], test[1])

		t.Run(tag, func(t *testing.T) {

			res := sum(test[0], test[1])

			if res != test[2] {
				t.Error("failed test")
			}

		})

	}

	t.Run("1+2", func(t *testing.T) {
		res := sum(1, 2)
		if res != 3 {
			t.Error("failed test")
		}
	})

	t.Run("2+3", func(t *testing.T) {
		res := sum(2, 3)
		if res != 7 {
			t.Error("failed test")
		}
	})

}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum(1, 2)
	}
}
