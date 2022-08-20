package main

import (
	"fmt"
	"sync"
	"time"
)

type Ch struct{
	sync.Mutex
}

type Philos struct{
	num, cnt int
	left, right *Ch
}

func (p Philos) eating(c chan *Philos, wg *sync.WaitGroup){
	for i := 0; i <= 2; i++{
		c <- &p
		if p.cnt <= 2{
			p.left.Lock()
			p.right.Lock()
			fmt.Println("Starting to eat ", p.num)
			p.cnt++
			fmt.Println("Finishing to eat ", p.num)
			p.left.Unlock()
			p.right.Unlock()
			wg.Done()
		}
	}
}

func allow(c chan *Philos, wg *sync.WaitGroup){
	for {
		if len(c) == 2{
			<- c
			<- c
			time.Sleep(20 * time.Millisecond)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	c := make(chan *Philos, 2)

	wg.Add(15)

	ch := make([] *Ch, 5)
	for i := 0; i < 5; i++{
		ch[i] = new(Ch)
	}
	philos := make([] *Philos, 5)
	for i := 0; i < 5; i++{
		philos[i] = &Philos{i+1,0,ch[i],ch[(i+1)%5]}
	}
	go allow(c, &wg)
	for i := 0; i < 5; i++{
		go philos[i].eating(c, &wg)
	}
	wg.Wait()
}