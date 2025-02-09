package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

const taskNum = 36

type Item struct {
	Name string
}

func main() {
	fmt.Printf("%d items\n", taskNum)
	fmt.Println()

	items := createItems(taskNum)

	validItems := make([]Item, 0, taskNum)
	var wg sync.WaitGroup
	var mu sync.Mutex
	for _, item := range items {
		wg.Add(1)
		go func(item Item) {
			defer wg.Done()
			if isValid := validate(item); isValid {
				mu.Lock()
				validItems = append(validItems, item)
				mu.Unlock()
			}
		}(item)
	}
	wg.Wait()

	fmt.Println()
	log.Printf("valid items: %d", len(validItems))
}

func createItems(taskNum int) []Item {
	items := make([]Item, taskNum)
	for i := 0; i < taskNum; i++ {
		items[i] = Item{Name: fmt.Sprintf("item%d", i+1)}
	}
	return items
}

func validate(item Item) bool {
	// something validation
	duration := rand.Intn(3) + 1
	time.Sleep(time.Duration(duration) * time.Second)

	log.Printf("%s: validation done.", item.Name)
	return true
}
