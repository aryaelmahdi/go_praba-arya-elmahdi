package main

import "fmt"
import "strconv"

func main(){
	var jumlah int
	var index int = 1
	var input string
	var err error

	fmt.Println("Masukkan Jumlah Tingkat")
	fmt.Scanln(&input)

	jumlah, err = strconv.Atoi(input)
	if err != nil {
		fmt.Println("Masukkan Angka")
		return
	}

	for a := 1; a<=jumlah ; a++ {
		for b := a; b<=jumlah ; b++ {
			fmt.Printf(" ")
		} 
		for c := 1; c<=index; c++ {
			if c%2 == 0 {
				fmt.Printf(" ")
			}else{
				fmt.Printf("*")
			}
		}
		fmt.Println()
		index += 2
	}
}