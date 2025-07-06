package main

import "fmt"

func multi(arr *[]int) {
	for i := 0; i < len(*arr); i++ {
		(*arr)[i] *= 2
	}
}

func main() {
	arr := []int{1, 2, 3}
	multi(&arr)
	fmt.Println(arr)
}
