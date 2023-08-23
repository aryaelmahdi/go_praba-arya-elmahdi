package main

import "fmt"


func prima(a int){
	if a > 1 {
		switch a {
			case 2:
				fmt.Printf("Bilangan Prima") 
				break
			case 3:
				fmt.Printf("Bilangan Prima") 
				break
			default :
				if a % 2 == 0 || a % 3 == 0 {
					fmt.Printf("Bukan Bilangan Prima")
					return
				}
				fmt.Printf("Bilangan Prima")
		}
		return
	}
	fmt.Printf("Bukan Bilangan Prima")
}

func main(){
	var bilangan int
	fmt.Scanln(&bilangan)
	prima(bilangan)
}