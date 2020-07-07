package main

import "fmt"

//need to be using a go routine for a channel to work
func channelDemo() {
	ch := make(chan string)
	//use a go routine or use this. this is saying use a channel to hold 1 value
	// ch := make(chan string, 1)

	go func(msg string) {
		//this puts the message into the channel
		ch <- msg
	}("hello")
	//this takes the message out of the channel
	recv := <-ch

	fmt.Println(recv)

}
