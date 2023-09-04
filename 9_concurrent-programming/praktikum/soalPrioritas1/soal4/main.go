package main

import (
	"fmt"
	"sync"
)

type Number struct {
	res int
	m   sync.Mutex
}

func (number *Number) factorial(num int) int {
	number.m.Lock()
	defer number.m.Unlock()
	number.res = 1
	for i := num; i > 0; i-- {
		number.res *= i
	}
	return number.res
}

func main() {
	var wg sync.WaitGroup
	number := Number{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		res := number.factorial(7)
		fmt.Println(res)
	}()
	wg.Wait()
}
