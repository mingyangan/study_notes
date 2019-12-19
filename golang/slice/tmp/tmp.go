package main

import (
	"fmt"
	"unsafe"
)

func change(s []int64) {
	s[0] = 100
}

func add(s []int64) {
	s = append(s, 200)
}

func main() {
	var slice = []int64{0, 1, 2}
	fmt.Printf("Init, slice len:%d, cap:%d, data:%v\n", len(slice), cap(slice), slice) // 3, 3, [0, 1, 2]
	change(slice)
	fmt.Printf("After Change Func, slice len:%d, cap:%d, data:%v\n", len(slice), cap(slice), slice) // 3, 3, [100, 1, 2]

	add(slice)
	fmt.Printf("After Add Func, slice len:%d, cap:%d, data:%v\n", len(slice), cap(slice), slice) // 3, 3, [100, 1, 2]

	slice = append(slice, 3)
	fmt.Printf("After Append, slice len:%d, cap:%d, data:%v\n", len(slice), cap(slice), slice) // 4, 6, [100, 1, 2, 3]
	add(slice)
	fmt.Printf("After Add Func, slice len:%d, cap:%d, data:%v\n", len(slice), cap(slice), slice) // 4, 6, [100, 1, 2, 3]
	fmt.Printf("Array Element[4]:%d\n",
		*(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&slice[0])) + uintptr(8*4)))) // 200
	return
}
