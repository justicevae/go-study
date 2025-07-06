package main

import (
	"fmt"
	"sync"
)

func addData(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 100; i++ {
		ch <- i
	}
	close(ch)
}

func printData(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range ch {
		fmt.Printf("接收数据: %d\n", i)
	}
}

func main() {
	ch := make(chan int, 10)

	var wg sync.WaitGroup
	wg.Add(2)

	go addData(ch, &wg)
	go printData(ch, &wg)

	wg.Wait()
}