package main

import (
	"fmt"
	"strconv"
	"time"
	"unsafe"
)

// Test3 string --> []byte
// 不使用只读区
func (t *testStruct) Test3() {
	var str = strconv.Itoa(int(time.Now().Nanosecond())) // str:="abc"这种显示的定义字符串，会被分配到只读区
	var buf []byte
	strPtr := (*[2]uintptr)(unsafe.Pointer(&str))
	buf = *(*[]byte)(unsafe.Pointer(&([3]uintptr{strPtr[0], strPtr[1], strPtr[1]})))
	fmt.Printf("Init String:%s\n", str)
	fmt.Printf("Init []Byte:%v, len:=%d\n", buf, len(buf))

	buf[0] = '0'
	fmt.Printf("After Change buf[0] = '0', []Byte Is:%s\n", buf)
	fmt.Printf("After Change buf[0] = '0', String Is:%s\n", str)
	return

}
