package main

import (
	"testing"
)

func Benchmark_calcWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calcWord()
	}

}
