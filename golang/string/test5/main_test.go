package main

import (
	"testing"
	"unsafe"
)

const (
	BufLen = 1024
)

func bytes2Str(buf []byte) (str string) {
	return *(*string)(unsafe.Pointer(&buf))
}

// Benchmark_TestCopyBytes 有copy的转化
func Benchmark_TestCopyBytes(b *testing.B) {

	var buf = make([]byte, BufLen)

	for i := 0; i < b.N; i++ {
		str := string(buf)
		_ = str
	}
}

// Benchmark_TestNoCopyBytes 无copy的转化
func Benchmark_TestNoCopyBytes(b *testing.B) {
	var buf = make([]byte, BufLen)

	for i := 0; i < b.N; i++ {
		str := bytes2Str(buf)
		_ = str
	}
}
