package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	taskNum := rand.Intn(10) + 1
	fmt.Printf("%d tasks\n", taskNum)

	var wg sync.WaitGroup
	wg.Add(taskNum)
	for i := 1; i <= taskNum; i++ {
		go task(&wg, i)
	}
	wg.Wait()
}

func task(wg *sync.WaitGroup, taskId int) {
	defer wg.Done()
	timeRequired := rand.Intn(10) + 1
	time.Sleep(time.Duration(timeRequired) * time.Second)
	log.Printf("task%d done.(%dsec)", taskId, timeRequired)
}
