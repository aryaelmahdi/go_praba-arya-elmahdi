// format Printf
// "%s" : string
// "%d" : int
// "%.[nilai]f" : float64

package main

import "fmt"

func pow(a float64, b int) float64 {
	var hasil float64 = 1
	if b >= 0 {
		for b > 0 {
			if b%2 == 1 {
				hasil *= a
			}
			a *= a
			b/=2
		}
	}else {
		for b < 0 {
			if b%2 == -1{
				hasil /= a
			}
			a *= a
			b/=2
		}
	}
	return hasil
}

func main(){

	fmt.Println(pow(2, 3)) // 8

	fmt.Println(pow(5, 3)) // 125

	fmt.Println(pow(10, 2)) // 100

	fmt.Println(pow(2, 5)) // 32

	fmt.Println(pow(7, 3)) // 343
}