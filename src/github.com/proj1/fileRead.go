package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readFileDemo() {
	// read1()
	// read2()
	read3()
}

func read1() {
	dat, _ := ioutil.ReadFile("dunkirk.txt")
	fmt.Println(dat)
	fmt.Println(string(dat))
}

func read2() {
	//f becomes a file handler class with a Read method.
	f, _ := os.Open("dunkirk.txt")
	b := make([]byte, 100)

	for {
		r, _ := f.Read(b)
		//When r == 0 that means we read 100 more bytes
		if r == 0 {
			break
		}
		// r will print the length of bytes
		fmt.Println(r)
		//string(b) will print the text
		fmt.Println(string(b))
	}
}

func read3() {
	f, _ := os.Open("dunkirk.txt")
	b := make([]byte, 100)

	for {
		r, err := f.Read(b)
		if err != nil && err != io.EOF {
			//Will throw an error if... there is an error
			panic(err)
		}
		if r == 0 {
			break
		}
		fmt.Println(string(b[:r]))
	}
}
