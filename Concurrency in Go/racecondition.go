package main

import (
	"fmt"
	"time"
)

/*
Write two goroutines which have a race condition
when executed concurrently. Explain what the
race condition is and how it can occur.
*/

/*
Races conditions occur due to the communication
between the two go routines since they use
the x variable at the same time which leads
to a non-deterministic output
*/

func add(x *int) {
	*x = *x + 1
}

func print(x int) {
	fmt.Println(x)
}

func main() {
	x := 1
	go add(&x)
	go print(x)

	time.Sleep(1 * time.Second)
	fmt.Println("Hello, playground")
}
