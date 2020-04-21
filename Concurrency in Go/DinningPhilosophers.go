package main

import (
	"fmt"
	"sync"
)

// ChopS type
type ChopS struct{ sync.Mutex }

// Philo type
type Philo struct {
	leftCS, rightCS *ChopS
}

func (p Philo) eat() {
	for {
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Println("“eating”")

		p.rightCS.Unlock()
		p.leftCS.Unlock()
	}
}

func dinningphilosophers() {
	CSticks := make([]*ChopS, 5)
	philos := make([]*Philo, 5)

	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	for i := 0; i < 5; i++ {
		philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5]}
	}

	for i := 0; i < 5; i++ {
		go philos[i].eat()
	}

}
