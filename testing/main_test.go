package main

import "testing"

func TestSum(t *testing.T) {
	res := sum(1, 2)
	if res != 3 {
		t.Error("failed test")
	}
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum(1, 2)
	}
}
