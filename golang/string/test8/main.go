package main

import (
	"bytes"
	"fmt"
	"strings"
)

func stringOperator(str []string) {
	var s string
	for i := 0; i < len(str); i++ {
		s += str[i]
	}
	_ = s
}

func stringSprintf(str []interface{}) {
	var s = fmt.Sprintf("%s", str...)
	_ = s
}

func stringJoin(str []string) {
	var s = strings.Join(str, "")
	_ = s
}

func stringBuffer(str []string) {
	var s string
	var buffer bytes.Buffer
	for i := 0; i < len(str); i++ {
		buffer.WriteString(str[i])
	}
	s = buffer.String()
	_ = s
}

func stringBuilder(str []string) {
	var s string
	var builder strings.Builder
	for i := 0; i < len(str); i++ {
		builder.WriteString(str[i])
	}
	s = builder.String()
	_ = s
}
