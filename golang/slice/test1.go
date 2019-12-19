package main

import "fmt"

// Test1 slice与数组的关系
func (t *testSlice) Test1() {
	var arr = [...]int{0, 1, 2}
	var slice = arr[0:2]
	fmt.Printf("Init, Array data:%v\n", arr)
	fmt.Printf("Init, slice data:%v\n", slice)
	println()

	arr[0] = 10
	fmt.Printf("Change arr[0] = 10, Array data:%v\n", arr)
	fmt.Printf("Change arr[0] = 10, slice data:%v\n", slice)
	println()

	slice[0] = 0
	fmt.Printf("Change slice[0] = 0, Array data:%v\n", arr)
	fmt.Printf("Change slice[0] = 0, slice data:%v\n", slice)
	println()

	slice = append(slice, 3)
	fmt.Printf("Slice Append 3, Array data:%v\n", arr)
	fmt.Printf("Slice Append 3, slice data:%v\n", slice)
	println()

	// append slice
	slice = append(slice, 4)
	fmt.Printf("Slice Append 4, Array data:%v\n", arr)
	fmt.Printf("Slice Append 4, slice data:%v\n", slice)
	println()

	slice[0] = 100
	fmt.Printf("Change slice[0] = 100, Array data:%v\n", arr)
	fmt.Printf("Change slice[0] = 100, slice data:%v\n", slice)
	println()
}
