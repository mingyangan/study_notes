package main

import (
	"testing"
)

var str []string
var inf []interface{}

const strCount = 100

func init() {
	for i := 0; i < strCount; i++ {
		str = append(str, "hello world")
		inf = append(inf, "hello world")
	}

}
func BenchmarkAddStringWithOperator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringOperator(str)
	}
}

func BenchmarkAddStringWithSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringSprintf(inf)
	}
}

func BenchmarkAddStringWithJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringJoin(str)
	}
}

func BenchmarkAddStringWithBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringBuffer(str)
	}
}
func BenchmarkAddStringWithBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringBuilder(str)
	}
}
