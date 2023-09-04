package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go multiples(ch)

	for result := range ch {
		fmt.Println(result)
	}
}

func multiples(ch chan int) {
	defer close(ch)
	for i := 1; i <= 10; i++ {
		ch <- 3 * i
	}
}
