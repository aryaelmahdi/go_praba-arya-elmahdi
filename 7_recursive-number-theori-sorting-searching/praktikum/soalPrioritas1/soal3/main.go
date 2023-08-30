package main

import "fmt"

func primeX(number int) int {
	var prime []int
	for i := 2; len(prime) < number; i++ {
		if i == 2 {
			prime = append(prime, i)
		}
		if i == 3 {
			prime = append(prime, i)
		}
		if i > 3 && i%2 != 0 && i%5 != 0 && i%7 != 0 {
			prime = append(prime, i)
		}
	}
	return prime[number-1]
}

func main() {

	fmt.Println(primeX(1)) // 2

	fmt.Println(primeX(5)) // 11

	fmt.Println(primeX(8)) // 19

	fmt.Println(primeX(9)) // 23

	fmt.Println(primeX(10)) // 29

}
