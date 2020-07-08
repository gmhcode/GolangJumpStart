package main

import (
	"fmt"
	"sync"
	"time"
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

func executeCountChannel() {
	c := make(chan string)
	go countGoChannel("sheep", c)
	msg := <- c
	fmt.Println(msg)
}

//how to make the computer to wait for your go routine
func executeCountWaitGroup() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		countGoWait("sheep")
		wg.Done()
	}()

	wg.Wait()
	// count2("fish")
}

func countGoWait(thing string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}

func countGoChannel(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}
}
