package main

import "fmt"

func tSubSlice() {
	subslice := memoryLeak()
	fmt.Println(subslice, "remainging underlying array: ", subslice[:cap(subslice)])
}

//this will have no memory leak, because we made a copy of the array
func noMemoryLeak() []int {

	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sub := make([]int, 3)
	copy(sub, s[1:4])
	return sub

}
func memoryLeak() []int {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//this is memory leak, because it's returning a slice, which has a reference to the array that was created. so the S array stays in memory
	return s[1:4]
}
