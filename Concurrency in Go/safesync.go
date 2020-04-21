package main

import (
	"fmt"
	"sync"
)

var mut sync.Mutex
var i int = 0
var wg sync.WaitGroup
var on sync.Once

// Iterating through a channel
// for i := range c { ..... }
// close(c)

// Select statement
// select {
//	case a = <- incha:
//		fmt.Println(a)
//	case outchan <- b:
//		fmt.Println(b)
//	case <- abort:
//		return
// }

// Once Synchronization +  Safe Concurrency
func setup() {
	fmt.Println("init")
}
func dostuff() {
	on.Do(setup)

	mut.Lock()
	i = i + 1
	mut.Unlock()

	fmt.Println("Hello")
	wg.Done()
}

func safesync() {
	wg.Add(2)
	go dostuff()
	go dostuff()
	wg.Wait()
}
