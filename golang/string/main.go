package main

import (
	"flag"
	"fmt"
	"reflect"
)

type testStruct struct{}

func println() {
	for i := 0; i < 50; i++ {
		fmt.Printf("*")
	}
	fmt.Printf("\n")
}
func main() {
	var funcName *string
	funcName = flag.String("func", "", "Function Name")
	flag.Parse()
	var t testStruct

	fs := reflect.ValueOf(&t)
	fv := fs.MethodByName(*funcName)
	if fv.Kind() == reflect.Func {
		fv.Call(nil)
	} else {
		fmt.Println("Not Function Name")
	}

}
