package main

import (
	"fmt"
	"sync"
)

func main() {
	text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua"
	result := frequent(text)

	for char, count := range result {
		fmt.Println(char, count)
	}
	fmt.Println(result)
}

func frequent(text string) map[string]int {
	result := make(map[string]int)
	var mu sync.Mutex
	var wg sync.WaitGroup

	char := []rune(text)
	charGo := len(char) / 5

	for i := 0; i < 5; i++ {
		wg.Add(1)
		start := i * charGo
		end := (i + 1) * charGo
		if i == 5-1 {
			end = len(char)
		}
		go func(start, end int) {
			defer wg.Done()
			localResult := make(map[string]int)
			for _, char := range char[start:end] {
				charStr := string(char)
				localResult[charStr]++
				mu.Lock()
				result[charStr]++
				newCount := result[charStr]
				mu.Unlock()
				fmt.Println(charStr, newCount)
			}
		}(start, end)
	}
	wg.Wait()
	return result
}
