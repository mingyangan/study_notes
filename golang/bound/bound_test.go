package main

import "testing"

func Benchmark_Testf1(b *testing.B) {
	is := make([]int, 256)
	bs := make([]byte, 102400)
	for i := 0; i < b.N; i++ {
		f1(is, bs)
	}

}
func Benchmark_Testf2(b *testing.B) {
	is := make([]int, 1024)
	bs := make([]byte, 1024)
	for i := 0; i < b.N; i++ {
		f2(is, bs)
	}

}
