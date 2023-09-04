package main

import (
	"fmt"
	"time"
)

func main() {
	go multiples(5)
	time.Sleep(3 * time.Second)
}

func multiples(num int) {
	for i := 0; i < num; i++ {
		fmt.Println(num * i)
	}
}
