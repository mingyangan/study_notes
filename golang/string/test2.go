package main

import (
	"fmt"
	"unsafe"
)

// string --> []byte zero copy
// 实现无copy的string转[]byte !!!
func (t *testStruct) Test2() {
	var str = "abc" // 显示的定义字符串，会被分配到只读区
	var buf []byte
	strPtr := (*[2]uintptr)(unsafe.Pointer(&str))
	buf = *(*[]byte)(unsafe.Pointer(&([3]uintptr{strPtr[0], strPtr[1], strPtr[1]})))
	fmt.Printf("Init String:%s\n", str)
	fmt.Printf("Init []Byte:%v, len:=%d\n", buf, len(buf))

	buf[0] = '0' // Be Careful!!
	fmt.Printf("After Change buf[0] = '0', []Byte Is:%s\n", buf)
	fmt.Printf("After Change buf[0] = '0', String Is:%s\n", str)

	/**(*byte)(unsafe.Pointer(strPtr[0])) = 'z'
	fmt.Printf("After Change String = 'z', []Byte Is:%s\n", buf)
	fmt.Printf("After Change String = 'z', String Is:%s\n", str)*/
	return

}
