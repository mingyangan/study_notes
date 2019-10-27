package main

import "fmt"

// Test5 测试copy
func (t *testSlice) Test5() {
	var sliceA = []int{1, 2, 3}

	var sliceB = []int{4, 5, 6}

	sliceC := make([]int, 1, 3)

	copy(sliceA[1:], sliceB)
	copy(sliceC, sliceB[0:3])

	fmt.Printf("sliceA len:%d, cap:%d, data:%v, addr:%p, elemAddr:%p\n",
		len(sliceA), cap(sliceA), sliceA, &sliceA, &sliceA[0])
	fmt.Printf("sliceB len:%d, cap:%d, data:%v, addr:%p, elemAddr:%p\n",
		len(sliceB), cap(sliceB), sliceB, &sliceB, &sliceB[0])
	fmt.Printf("sliceC len:%d, cap:%d, data:%v, addr:%p, elemAddr:%p\n",
		len(sliceC), cap(sliceC), sliceC, &sliceC, &sliceC[0])

	sliceB[0] = 100
	fmt.Printf("sliceA len:%d, cap:%d, data:%v, addr:%p, elemAddr:%p\n",
		len(sliceA), cap(sliceA), sliceA, &sliceA, &sliceA[0])
	fmt.Printf("sliceB len:%d, cap:%d, data:%v, addr:%p, elemAddr:%p\n",
		len(sliceB), cap(sliceB), sliceB, &sliceB, &sliceB[0])
	fmt.Printf("sliceC len:%d, cap:%d, data:%v, addr:%p, elemAddr:%p\n",
		len(sliceC), cap(sliceC), sliceC, &sliceC, &sliceC[0])

}
