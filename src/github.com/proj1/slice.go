package main

import "fmt"

func slice1() {
	x := make([]int, 0)

	x = append(x, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	fmt.Println(x)
}

func slice2() {
	x := make([]int, 3, 10)

	fmt.Println(x)

	x[1] = 2
	fmt.Println(x)

	x = append(x, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14)
	fmt.Println(x)

	fmt.Println("x -> ", x)
	fmt.Println("x[2:5] -> ", x[2:5])
	fmt.Println("x[:5] -> ", x[:5])
	fmt.Println("x[:] -> ", x[:])
	fmt.Println("x[:5] -> ", x[:5])
	fmt.Println("x[3:] -> ", x[3:])

}

func slice3() {
	s := make([]string, 3)
	fmt.Println("String-Slice", s)
	s[0] = "0"
	s[2] = "0"

	s = append(s, "c", "d")
	fmt.Println(s)
}
