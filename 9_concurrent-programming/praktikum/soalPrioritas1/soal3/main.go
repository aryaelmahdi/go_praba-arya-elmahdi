package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 10)

	go multiples(ch)

	for result := range ch {
		fmt.Println(result)
	}
	time.Sleep(1 * time.Second)
}

func multiples(ch chan int) {
	defer close(ch)
	for i := 1; i <= 10; i++ {
		ch <- 3 * i
	}
}
