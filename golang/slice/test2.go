package main

import "fmt"

// Test2 append
func (t *testSlice) Test2() {
	var slice []int
	fmt.Printf("Init, slice len:%d, cap:%d, data:%v\n", len(slice), cap(slice), slice)

	slice = append(slice, 0)
	fmt.Printf("After Append 0, slice len:%d, cap:%d, data:%v\n", len(slice), cap(slice), slice)

	slice = append(slice, 1)
	fmt.Printf("After Append 1, slice len:%d, cap:%d, data:%v\n", len(slice), cap(slice), slice)
	slice = append(slice, 2)
	fmt.Printf("After Append 2, slice len:%d, cap:%d, data:%v\n", len(slice), cap(slice), slice)

	slice = append(slice, 3)
	slice = append(slice, 4)
	fmt.Printf("After Append 3,4, slice len:%d, cap:%d, data:%v\n", len(slice), cap(slice), slice)

	slice = append(slice, 5)
	fmt.Printf("After Append 5, slice len:%d, cap:%d, data:%v\n", len(slice), cap(slice), slice)
}
