package main

import (
	"fmt"
	"sort"
)

func sortdemo() {
	s := []string{"z", "x", "b", "a", "y"}
	sort.Strings(s)
	fmt.Println("Sorted String: ", s)

	nums := []int{7, 5, 4, 1, 2}
	sort.Ints(nums)
	fmt.Println("Sorted Numbers: ", nums)
}
