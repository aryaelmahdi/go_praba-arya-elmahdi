package main

import "fmt"
import "strconv"

func main(){
	var angka int
	var input string
	var err error

	fmt.Println("Pengecekan Ganjil Genap \n")
	fmt.Println("Masukkan Angka : ")

	fmt.Scanln(&input)
	angka, err = strconv.Atoi(input)

	if err != nil{
		fmt.Println("nilai harus angka")
		return
	}
	checkNum(angka)
}

func checkNum(a int){
	hasil := a%2
	switch hasil{
	case 0 :
		fmt.Println(a, "adalah angka genap")
		break
	case 1 : 
		fmt.Println(a, "adalah angka ganjil")
		break
	}
}