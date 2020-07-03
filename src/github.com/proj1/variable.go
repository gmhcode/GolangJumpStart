package main

func add(a, b int) int {
	return a + b
}

func Boolean() bool {
	var b bool
	b = true
	return b
}

func showvar() {
	var a string = "mystring"
	println(a)

	var e int
	println(e)

	var b, c int = 1, 2
	println(b, c)

	d := "love go"

	f := a + " : " + d
	println(f)
}
