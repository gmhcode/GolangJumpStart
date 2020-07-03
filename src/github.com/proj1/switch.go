package main

func whatday(n int) {
	if n != 0 && n <= 7 {
		switch n {
		case 1:
			println("Monday")
		case 2:
			println("Tuesday")
		case 6:
			println("Saturday")
		case 7:
			println("Hola - Sunday")
		default:
			println("it is a weekday")
		}
	} else {
		println("Wrong input 1-7 accepted")
	}
}
