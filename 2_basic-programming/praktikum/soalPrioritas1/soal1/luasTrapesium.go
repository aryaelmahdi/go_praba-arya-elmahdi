package main

import "fmt"
import "strconv"

func main(){
	var atas, bawah, tinggi int
	var input string
	var err error

	fmt.Println("\n Menghitung luas trapesium \n")
	fmt.Println("Masukkan nilai sisi Atas")
	fmt.Scanln(&input)
	atas, err = strconv.Atoi(input) 
	if err != nil {
		fmt.Println("nilai harus angka")
		return
	}

	fmt.Println("Masukkan nilai sisi Bawah")
	fmt.Scanln(&input)
	bawah, err = strconv.Atoi(input)
	if err != nil {
		fmt.Println("nilai harus angka")
		return
	}

	fmt.Println("Masukkan nilai Tinggi")
	fmt.Scanln(&input)
	tinggi, err = strconv.Atoi(input)
	if err != nil {
		fmt.Println("nilai harus angka")
		return
	}

	calculate(atas, bawah, tinggi)
}

func calculate(a,b,c int){
	luas := (float64((a+b)*c)/2)
	fmt.Println("Luas trapesium = ", luas)
}
