package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

var taskNum = rand.Intn(10) + 1

func main() {
	fmt.Printf("%d tasks\n", taskNum)

	func1(taskNum)
	func2(taskNum)
	func3(taskNum)
}

func func1(taskNum int) {
	var wg sync.WaitGroup
	wg.Add(taskNum)
	for i := 1; i <= taskNum; i++ {
		go task(&wg, "func1", i)
	}
	wg.Wait()
}

func func2(taskNum int) {
	var wg sync.WaitGroup
	for i := 1; i <= taskNum; i++ {
		wg.Add(1)
		go task(&wg, "func2", i)
	}
	wg.Wait()
}

func task(wg *sync.WaitGroup, funcName string, taskId int) {
	defer wg.Done()
	duration := rand.Intn(3) + 1
	time.Sleep(time.Duration(duration) * time.Second)
	log.Printf("%s: task%d done.(%dsec)", funcName, taskId, duration)
}

func func3(taskNum int) {
	var wg sync.WaitGroup
	for i := 1; i <= taskNum; i++ {
		wg.Add(1)
		go func(funcName string, taskId int) {
			duration := rand.Intn(3) + 1
			time.Sleep(time.Duration(duration) * time.Second)
			log.Printf("%s: task%d done.(%dsec)", funcName, taskId, duration)
			defer wg.Done()
		}(string("func3"), i)
	}
	wg.Wait()
}
