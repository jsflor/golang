package main

import (
	"fmt"
)

func prod(x int, y int, c chan int) {
	c <- x * y
}

func channels() {
	// c := make(chan int, 3) buffered
	c := make(chan int) // unbuffered

	go prod(1, 2, c)
	go prod(3, 4, c)

	a := <-c
	b := <-c

	fmt.Printf("%v", a*b)
}
