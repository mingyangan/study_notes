package main

import "testing"

var (
	arr    = make([]TestObj, 1000)
	bigArr = make([]TestBigObj, 1000)
)

func RangeItemTestObj(arr []TestObj) {
	for _, item := range arr {
		_ = item
	}
}

func RangeIndexTestObj(arr []TestObj) {
	for index := range arr {
		_ = arr[index]
	}
}
func RangeItemTestBigObj(arr []TestBigObj) {
	for _, item := range arr {
		_ = item
	}
}

func RangeIndexTestBigObj(arr []TestBigObj) {
	for index := range arr {
		_ = arr[index]
	}
}

func Benchmark_TestRangeItemTestObj(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RangeItemTestObj(arr)

	}

}
func Benchmark_TestRangeIndexTestObj(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RangeIndexTestObj(arr)
	}

}

func Benchmark_TestRangeItemTestBigObj(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RangeItemTestBigObj(bigArr)

	}

}
func Benchmark_TestRangeIndexTestBigObj(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RangeIndexTestBigObj(bigArr)
	}

}
