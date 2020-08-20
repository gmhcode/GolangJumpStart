package main

func pointer1() {
	var a *int
	b := 6
	a = &b

	println("Address ", a)
	println("value ", *a)
}

func pointer2() {

	c := 5
	// increment(c)
	// increment(c)
	// println("called from pointer 2 func :", c)

	incrementbyptr(&c)
	increment(c)
	println("c: ", c)
	println("called from pointer 2 func :", c)
}

func increment(c int) {
	c++
	println("Called within increment", c)
}

func incrementbyptr(c *int) {
	*c++
	println("Called within increment", *c)
}
