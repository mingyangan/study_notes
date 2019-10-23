package main

import "testing"

func Benchmark_change(b *testing.B) {
	slice := make([]int, 100)
	for i := 0; i < b.N; i++ {
		change(slice)

	}
}

func Benchmark_changeSlicePtr(b *testing.B) {
	slice := make([]int, 100)
	for i := 0; i < b.N; i++ {
		changeSlicePtr(&slice)
	}
}

func Benchmark_changeElemPtr(b *testing.B) {
	slice := make([]*int, 100)
	for i := 0; i < b.N; i++ {
		changeElemPtr(slice)
	}
}
