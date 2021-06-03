package main

import (
	"fmt"
	"sync"
	"time"
)

type Chops struct {
	sync.Mutex
}
type Philo struct {
	left *Chops
	right *Chops
	eateanFood int
	id int
}
func (p *Philo) Eat(host *chan bool, wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		*host <- true
		p.left.Lock()
		p.right.Lock()
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("starting to eat %d\n", p.id)
		p.eateanFood++
		p.left.Unlock()
		p.right.Unlock()
		<- *host
	}
	wg.Done()
}
func main() {
	CSticks := make([]*Chops, 5)
	host := make(chan bool, 2)
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		CSticks[i] = new(Chops)
	}
	philos := make([]*Philo, 5) 
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{left: CSticks[i], right: CSticks[(i + 1) % 5], eateanFood: 0, id: i + 1}
	}


	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philos[i].Eat(&host, &wg)
	}
	wg.Wait()
	

	for _, value := range philos {
		fmt.Printf("id : %d  eat: %d\n", value.id,  value.eateanFood)
	}
}