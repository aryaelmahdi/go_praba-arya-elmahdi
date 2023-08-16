// Nilai 80 - 100: A
// Nilai 65 - 79: B
// Nilai 50 - 64: C
// Nilai 35 - 49: D
// Nilai 0 - 34: E
package main

import "fmt"
import "strconv"

func main(){
	var nilai int
	var input string
	var err error

	fmt.Println("Masukkan nilai")
	fmt.Scanln(&input)

	nilai, err = strconv.Atoi(input)
	if err != nil {
		fmt.Println("Nilai tidak valid")
		return
	}

	if nilai >=0 && nilai <=34 {
		fmt.Println("Nilai ", nilai, " : E")
	}else if nilai >34 && nilai <= 49 {
		fmt.Println("Nilai ", nilai, " : D")
	}else if nilai >49 && nilai <= 64 {
		fmt.Println("Nilai ", nilai, " : C")
	}else if nilai >64 && nilai <= 79 {
		fmt.Println("Nilai ", nilai, " : B")
	}else if nilai >79 && nilai <=100 {
		fmt.Println("Nilai ", nilai, " : A")
	}else {
		fmt.Println("Nilai ", nilai, " : Tidak Valid")
	}
}