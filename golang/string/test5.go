package main

import (
	"fmt"
	"unsafe"
)

// Test5 []byte --> string
// 实现无copy的[]byte转string
func (t *testStruct) Test5() {
	buf := []byte{'a', 'b', 'c'}

	str := *(*string)(unsafe.Pointer(&buf))
	fmt.Printf("Init, []Byte Is:%s\n", buf)
	fmt.Printf("Init, String Is:%s\n", str)

	buf[0] = '0'
	fmt.Printf("After Change buf[0] = '0', []Byte Is:%s\n", buf)
	fmt.Printf("After Change buf[0] = '0', String Is:%s\n", str)

}
