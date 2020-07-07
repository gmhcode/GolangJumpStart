package main

import "fmt"

func advancedSlice() {
	s := make([]int, 3)
	for i := 0; i < 3; i++ {
		println("i value", i)
		s[i] = i + 1
	}
	s = append(s, 4, 6, 7, 8, 9)

	fmt.Println("s = ", s)

	//copy Slice

	// b := make([]int, len(s))
	// copy(b, s)
	// fmt.Println("b = ", b)

	//cut
	//this will return an array of everything in the slice up to but not including index 2 and index 4 and onward. so it does not return index 2 and 3
	// s = append(s[:2], s[4:]...)
	// fmt.Println("cut(3,4) s = ", s)

	//Del by index
	s = delbyindex(2, s)
	fmt.Println("s = ", s)

}

func delbyindex(i int, a []int) []int {
	//we return everything up to i and everything i + 1 and after. so we dont return i
	a = append(a[:i], a[i+1:]...)
	return a
}
