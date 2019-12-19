package main

import (
	"fmt"
	"unicode"
)

func main() {
	var r byte = 0x56
	fmt.Println(unicode.IsLetter(rune(r)))
}
