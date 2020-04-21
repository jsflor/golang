package main

import (
	"fmt"
	"sync"
)

func foo(wg *sync.WaitGroup) {
	fmt.Println("Go routine!")
	wg.Done()
}

func syncWG() {
	var wg sync.WaitGroup

	wg.Add(1)
	go foo(&wg)
	wg.Wait()

	fmt.Println("Hello, playground")
}
