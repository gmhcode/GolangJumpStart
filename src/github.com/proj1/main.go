package main

import "fmt"

func main() {
	// fmt.Println("Hello World")
	// c := add(5, 6)
	// println(c)
	// z := Boolean()
	// println(z)
	// showvar()
	// loop1()
	// loop2()
	// loop3()
	// loop4()
	// loop5()
	// oddeven()
	// days()
	// whatday(4)
	// arr1()
	// arr2()
	// matrix()
	// imatrix(3)
	// slice1()
	// slice2()
	// slice3()
	// pointer1()
	// pointer2()
	// variadic()
	// recursion()
	// recursionPattern(5)
	// sortdemo()
	// structDemo()
	// advancedSlice()
	// readFileDemo()
	// writeFileDemo()
	// mapDemo()
	// methodDemo()
	// interfaceDemo()
	// goRoutine()
	// channelDemo()
	// executeCountWaitGroup()
	// executeCountChannel()

	// s := make(Set)
	// //The extra empty bracket on the struct means that the struct is empty
	// s["item1"] = struct{}{}
	// s["item2"] = struct{}{}
	// fmt.Println(getValues(s))
	// sameUnderlyingSliceArray()
	copyFunction()

}

// copy makes new memory for the array its copying, so it wont change the original array
func copyFunction() {
	s1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("s1: ", s1)
	s2 := make([]int, 2)
	n := copy(s2, s1[2:4])
	fmt.Println("s1[2:4]", s1[2:4])
	fmt.Println("Number of items copied: ", n)
	//s2 is the slice of the s1 array in memory, so changing it will effect the s1 array
	s2[0] = 10
	fmt.Println("s1: ", s1)
	fmt.Println("s2: ", s2)
}

func sameUnderlyingSliceArray() {
	s1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("s1: ", s1)
	s2 := s1[2:4]
	//s2 is the slice of the s1 array in memory, so changing it will effect the s1 array
	s2[0] = 10
	fmt.Println("s1: ", s1)

}

func slices() {
	a := []int{1, 2, 3, 4, 5, 6}
	b := a[2:4]
	c := a[:3]
	d := a[3:]

	fmt.Println("Capacity of b: ", cap(b))
	// Prints Capacity of b:  4

	fmt.Println("Slices a", a, " b:", b, " c: ", c, " d: ", d)
	//Prints Slices a [1 2 3 4 5 6]  b: [3 4]  c:  [1 2 3]  d:  [4 5 6]

	fmt.Println("What b actually sees: ", b[:cap(b)])
	//Prints What b actually sees:  [3 4 5 6]
	//because be started at [2] so it can see the whole slice
}
