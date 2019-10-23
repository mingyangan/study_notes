package main

import "fmt"

func (t *testRange) Test1() {
	arr := [...]TestObj{
		{
			Name: "a",
		},
		{
			Name: "b",
		},
		{
			Name: "c",
		},
	}
	var slice []*TestObj

	for _, item := range arr {
		if item.Name != "b" {
			slice = append(slice, &item)
		}
	}
	fmt.Println(slice)

	// fmt.Printf("&arr[0]=%p\n", &arr[0])
	// fmt.Printf("&arr[1]=%p\n", &arr[1])
	// fmt.Printf("&arr[2]=%p\n", &arr[2])
}
