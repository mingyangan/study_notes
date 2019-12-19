package main

import (
	"testing"
	"unsafe"
)

const (
	StringLen = 1024
)

func str2Bytes(str string) (buf []byte) {
	strPtr := (*[2]uintptr)(unsafe.Pointer(&str))
	buf = *(*[]byte)(unsafe.Pointer(&([3]uintptr{strPtr[0], strPtr[1], strPtr[1]})))
	return
}

// Benchmark_TestCopy 有copy的转化
func Benchmark_TestCopy(b *testing.B) {

	var str string
	for i := 0; i < StringLen; i++ {
		str += "a"
	}

	for i := 0; i < b.N; i++ {
		buf := []byte(str)
		_ = buf
	}
}

// Benchmark_TestNoCopy 无copy的转化
func Benchmark_TestNoCopy(b *testing.B) {
	var str string
	for i := 0; i < StringLen; i++ {
		str += "a"
	}

	for i := 0; i < b.N; i++ {
		buf := str2Bytes(str)
		_ = buf
	}
}
