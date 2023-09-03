package main

import "fmt"

var (
	fibo = map[int]int{}
)

func fibonacci(num int) int {
	if value, found := fibo[num]; found {
		return value
	}

	if num <= 1 {
		fibo[num] = num
		return num
	}

	res := fibonacci(num-2) + fibonacci(num-1)
	fibo[num] = res
	return res
}

func main() {
	fmt.Println(fibonacci(0)) // 0

	fmt.Println(fibonacci(2)) // 1

	fmt.Println(fibonacci(9)) // 34

	fmt.Println(fibonacci(10)) // 55

	fmt.Println(fibonacci(12)) // 144
}
