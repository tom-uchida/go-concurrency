package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Item struct {
	Name string
}

func main() {
	const itemNum = 100
	fmt.Printf("%d items\n", itemNum)
	fmt.Println()

	items := createItems(itemNum)

	validItems := make([]Item, 0, itemNum)
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

func createItems(itemNum int) []Item {
	items := make([]Item, itemNum)
	for i := 0; i < itemNum; i++ {
		items[i] = Item{Name: fmt.Sprintf("item%d", i+1)}
	}
	return items
}
func validate(item Item) bool {
	// something validation
	time.Sleep(time.Second * 1)

	log.Printf("%s: validation done.", item.Name)
	return true
}
