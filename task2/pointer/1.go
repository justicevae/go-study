package main

import "fmt"

func add(num *int) {
	*num += 10
}

func main() {
	num := 1
	add(&num)
	fmt.Println("修改后的值:", num)
}
