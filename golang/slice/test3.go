package main

import (
	"fmt"
)

// 扩容地址
func (t *testSlice) Test3() {
	var slice = []int{0}
	fmt.Printf("Init, slice len:%d, cap:%d, data:%v, &slice=%p, &slice[0]=%p\n",
		len(slice), cap(slice), slice, &slice, &slice[0])

	slice = append(slice, 1)
	fmt.Printf("Append 1, slice len:%d, cap:%d, data:%v, &slice=%p, &slice[0]=%p\n",
		len(slice), cap(slice), slice, &slice, &slice[0])
}
