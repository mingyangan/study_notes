package main

import (
	"fmt"
	"unsafe"
)

var (
	a = 100
)

func change(s []int) {
	s[0] = 100
}

func add(s []int) {
	s = append(s, 200)
}

func changeSlicePtr(s *[]int) {
	(*s)[0] = 100
}

func changeElemPtr(s []*int) {
	s[0] = &a
}

// Test4 slice作为函数参数
func (t *testSlice) Test4() {
	var slice = []int{0, 1, 2}
	fmt.Printf("Init, slice len:%d, cap:%d, data:%v\n", len(slice), cap(slice), slice)
	// 函数内赋值
	change(slice)
	fmt.Printf("After Change Func, slice len:%d, cap:%d, data:%v\n", len(slice), cap(slice), slice)
	println()

	// 函数内append
	add(slice)
	fmt.Printf("After Add Func, slice len:%d, cap:%d, data:%v\n", len(slice), cap(slice), slice)
	println()

	// 上面可能是因为扩容，所以这里先扩容下，再传到函数append
	slice = append(slice, 3)
	fmt.Printf("After Append, slice len:%d, cap:%d, data:%v\n", len(slice), cap(slice), slice)
	add(slice)
	fmt.Printf("After Add Func, slice len:%d, cap:%d, data:%v\n", len(slice), cap(slice), slice)
	// unsafe 看一下,底层数组其实已经赋值了
	fmt.Printf("Array Element[4]:%d\n", *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&slice[0])) + uintptr(8*4))))
	println()
}
