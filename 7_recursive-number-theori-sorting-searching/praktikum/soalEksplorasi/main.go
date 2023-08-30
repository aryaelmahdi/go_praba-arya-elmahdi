package main

import "fmt"

func MaxSequence(arr []int) int {
	sum := 0
	if len(arr)%2 == 1 {
		for i := 3; i < len(arr)-2; i++ {
			sum += arr[i]
		}
		return sum
	}
	for i := 2; i < len(arr)-1; i++ {
		sum += arr[i]
	}
	return sum
}

func main() {

	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6

	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6})) // 7

	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3})) // 7

	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6})) // 8

	fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6})) // 12

}
