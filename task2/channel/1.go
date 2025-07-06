package main

import (
	"fmt"
	"time"
)

func add(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

func print(ch <-chan int) {
	for i := range ch {
		fmt.Printf("接收数据: %d\n", i)
	}
}

func main() {
	ch := make(chan int)
	go add(ch)
	go print(ch)

	time.Sleep(2 * time.Second)
}
