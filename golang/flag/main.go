package main

import (
	"flag"
	"fmt"
)

var (
	x = "var x"
)

func init() {
	fmt.Println(x)
	str := flag.String("config", "abc", "default")

	flag.Parse()

	i := flag.Int("type", 1, "default type")
	flag.Parse()

	fmt.Println(*str)
	fmt.Println(*i)
}

func main() {

}
