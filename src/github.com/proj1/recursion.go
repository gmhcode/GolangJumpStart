package main

import "fmt"

func recursion() {
	z := fact(5)
	// println(z)
	fmt.Printf("Factorial : %d\n", z)
}

func fact(n int) int {
	if n == 1 {
		return 1
	}
	return n * fact(n-1)
}

func recursionPattern(num int) {
	for i := num; i > 0; i-- {
		fmt.Print("*")
	}
	fmt.Println()
	if num > 0 {
		recursionPattern(num - 1)
	}
}
