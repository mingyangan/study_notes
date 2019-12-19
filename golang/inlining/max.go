package main

import "fmt"

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func F() int {
	const a = 100
	const b = 20
	m := Max(a, b)
	return m
}

func main() {
	fmt.Println(F())
}
