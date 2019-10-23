package main

import "testing"

func AppendSliceWithMake() {
	a := make([]int, 100)
	for i := 0; i < len(a); i++ {
		a = append(a)
	}
}
func AppendSliceWithoutMake() {
	var a []int
	var length = 100
	for i := 0; i < length; i++ {
		a = append(a, i)
	}
}

// slice append make
func Benchmark_SliceMake(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AppendSliceWithMake()
	}
}

// slice append without make
func Benchmark_SliceWithoutMake(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AppendSliceWithoutMake()
	}
}
