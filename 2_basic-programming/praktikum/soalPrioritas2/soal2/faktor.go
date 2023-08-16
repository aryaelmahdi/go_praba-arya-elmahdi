package main

import "fmt"
import "strconv"

func main(){
	var angka int
	var input string
	var err error

	fmt.Println("Masukkan Angka : ")
	fmt.Scanln(&input)

	angka, err = strconv.Atoi(input)
	if err != nil {
		fmt.Println("Nilai input tidak sesuai")
		return
	}

	for i:=1 ; i<=angka ; i++ {
		if angka%i == 0 {
			fmt.Println(i)
		}
	}
}