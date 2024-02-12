package main

import (
	"fmt"
	"sync"
)

func main() {
	ps := NewPubSub[string]()
	wg := sync.WaitGroup{}

	sub1 := ps.Subscribe()
	wg.Add(1)
	go subscriber(&wg, sub1, 1)

	sub2 := ps.Subscribe()
	wg.Add(1)
	go subscriber(&wg, sub2, 2)

	ps.Publish("one")
	ps.Publish("two")
	ps.Publish("three")
	ps.Close()

	wg.Wait()
	fmt.Println("completed.")
}

func subscriber(wg *sync.WaitGroup, sub <-chan string, id int) {
	for {
		val, ok := <-sub
		if !ok {
			fmt.Printf("sub %d, exiting\n", id)
			wg.Done()
			return
		}
		fmt.Printf("sub %d, value %s\n", id, val)
	}
}
