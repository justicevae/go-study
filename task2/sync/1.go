package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	mu    sync.Mutex
	count int
}

func (c *SafeCounter) Increment() {
	defer c.mu.Unlock()
	c.mu.Lock()
	c.count++
}

func main() {
	counter := SafeCounter{}
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
		}()
	}

	wg.Wait()
	fmt.Printf("计数器的值: %d\n", counter.count)
}
