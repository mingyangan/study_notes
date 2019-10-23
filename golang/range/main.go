package main

import (
	"flag"
	"fmt"
	"reflect"
)

type TestObj struct {
	Name string
}

type TestBigObj struct {
	Name string
	i1   int64
	i2   int64
	i3   int64
	i4   int64
	i5   int64
	i6   int64
	i7   int64
	i8   int64
	i9   int64
	i10  int64
	i11  int64
	i12  int64
	i13  int64
	i14  int64
	i15  int64
	i16  int64
	i17  int64
	i18  int64
}

type testRange struct{}

func main() {
	var funcName *string
	funcName = flag.String("func", "", "Function Name")
	flag.Parse()
	var t testRange

	fs := reflect.ValueOf(&t)
	fv := fs.MethodByName(*funcName)
	if fv.Kind() == reflect.Func {
		fv.Call(nil)
	} else {
		fmt.Println("Not Function Name")
	}
}
