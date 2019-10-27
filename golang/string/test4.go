package main

import (
	"fmt"
	"unsafe"
)

// []byte --> string
// 验证[]byte转string是有copy的
func (t *testStruct) Test4() {
	buf := []byte{'a', 'b', 'c'}
	str := string(buf)

	fmt.Printf("Init, []Byte Is:%s\n", buf)
	fmt.Printf("Init, String Is:%s\n", str)

	buf[0] = '0'
	fmt.Printf("After Change buf[0] = '0', []Byte Is:%s\n", buf)
	fmt.Printf("After Change buf[0] = '0', String Is:%s\n", str)

	strPtr := (*[2]uintptr)(unsafe.Pointer(&str))      // 将string的两个属性转化为uintptr数组的指针
	fmt.Printf("String Data Address:%#x\n", strPtr[0]) // (*strPtr)[0]，string中指向数据的指针的值，即数据的地址
	fmt.Printf("Byte Address:%p\n", &buf[0])

}
