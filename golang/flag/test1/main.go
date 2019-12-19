package main

import (
	"flag"
	"fmt"
)

func init() {
	str := flag.String("config", "", "default config")
	flag.Parse()

	fmt.Printf("str:=%s\n", *str)

	
	flag.Parse()
	fmt.Printf("str:=%s\n", *str)
}
func main() {

}
