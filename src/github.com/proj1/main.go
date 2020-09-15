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

	s := make(Set)
	//The extra empty bracket on the struct means that the struct is empty
	s["item1"] = struct{}{}
	s["item2"] = struct{}{}
	fmt.Println(getValues(s))
}
