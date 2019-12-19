package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

var strA = "abc"
var strB = "efg"

func BenchmarkAddStringWithOperator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strA + "," + strB
	}
}

func BenchmarkAddStringWithSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s,%s", strA, strB)
	}
}

func BenchmarkAddStringWithJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strings.Join([]string{strA, strB}, ",")
	}
}

func BenchmarkAddStringWithBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var buffer bytes.Buffer
		buffer.WriteString(strA)
		buffer.WriteString(",")
		buffer.WriteString(strB)
		_ = buffer.String()
	}
}
func BenchmarkAddStringWithBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		builder.WriteString(strA)
		builder.WriteString(",")
		builder.WriteString(strB)
		_ = builder.String()
	}
}
