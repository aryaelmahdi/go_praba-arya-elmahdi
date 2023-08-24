package main

import "fmt"


func primeNumber(a int) bool{
	isPrime := false
	if a > 1 {
		switch a {
			case 2:
				// fmt.Printf("Bilangan Prima") 
				isPrime = true
				break
			case 3:
				// fmt.Printf("Bilangan Prima") 
				isPrime = true
				break
			default :
				if a % 2 == 0 || a % 3 == 0 || a % 7 == 0 {
					// fmt.Printf("Bukan Bilangan Prima")
					break
				}
				isPrime = true
				// fmt.Printf("Bilangan Prima")
		}
	}
	return isPrime
	// fmt.Printf("Bukan Bilangan Prima")
}

// func main(){
// 	var bilangan int
// 	fmt.Scanln(&bilangan)
// 	prima(bilangan)
// }

func main() {

	fmt.Println(primeNumber(1000000007)) // true
	
	fmt.Println(primeNumber(13)) // true
	
	fmt.Println(primeNumber(17)) // true
	
	fmt.Println(primeNumber(20)) // false
	
	fmt.Println(primeNumber(35)) // false
	
	}