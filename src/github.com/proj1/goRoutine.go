package main

import (
	"fmt"
)

func goRoutine() {
	// count(5)
	go count(5)
	// time.Sleep(time.Second * 1)

}

func count(num int) {
	for i := 0; i <= num; i++ {
		fmt.Println(i)
	}
}
