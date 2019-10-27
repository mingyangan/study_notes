package main

import (
	"fmt"
	"strconv"
	"time"
	"unsafe"
)

func getStringDataAddr(str *string) uintptr {
	return ((*[2]uintptr)(unsafe.Pointer(str)))[0]
}

// Test6 string数据的复用
// 在只读区的(显示定义)的相同字符串，共享数据
// 赋值的字符串，共享数据（哪怕有截取）
// 非只读区的，不共享数据
// []byte相同数据不共享
func (t *testStruct) Test6() {
	// 只读内存区
	var str1 = "abc"
	var str2 = "abc"

	fmt.Printf("Init, Str1:%s, Addr:%p, Data Addr:%#x\n", str1, &str1, getStringDataAddr(&str1))
	fmt.Printf("Init, Str2:%s, Addr:%p, Data Addr:%#x\n", str2, &str2, getStringDataAddr(&str2))
	fmt.Println()

	var str3 = strconv.Itoa(time.Now().Nanosecond())[0:3] // 只是为了获得一个不在只读区的字符串
	var str4 = str3[0:1]
	var str5 = strconv.Itoa(time.Now().Nanosecond())[0:3]
	fmt.Printf("Init, Str3:%s, Addr:%p, Data Addr:%#x\n", str3, &str3, getStringDataAddr(&str3))
	fmt.Printf("Init, Str4:%s, Addr:%p, Data Addr:%#x\n", str4, &str4, getStringDataAddr(&str4))
	fmt.Printf("Init, Str5:%s, Addr:%p, Data Addr:%#x\n", str5, &str5, getStringDataAddr(&str5))
	fmt.Println()

	*(*byte)(unsafe.Pointer(((*[2]uintptr)(unsafe.Pointer(&str3)))[0])) = '0'
	fmt.Printf("After Change, Str3:%s, Addr:%p, Data Addr:%#x\n", str3, &str3, getStringDataAddr(&str3))
	fmt.Printf("After Change, Str4:%s, Addr:%p, Data Addr:%#x\n", str4, &str4, getStringDataAddr(&str4))
	fmt.Printf("After Change, Str5:%s Addr:%p, Data Addr:%#x\n", str5, &str5, getStringDataAddr(&str5))
	fmt.Println()

	buf1 := []byte{'1', '2'}
	buf2 := []byte{'1', '2'}
	fmt.Printf("Init Buf1:%v, addr:%p\n", buf1, &buf1[0])
	fmt.Printf("Init Buf2:%v, addr:%p\n", buf2, &buf2[0])
}
