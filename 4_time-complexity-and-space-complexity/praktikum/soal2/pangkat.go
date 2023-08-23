// format Printf
// "%s" : string
// "%d" : int
// "%.[nilai]f" : float64

package main

import "fmt"

func pangkat1(a float64, b int) float64 {
	var hasil float64 = 1
	if b >= 0 {
		// for i := 1.0; i <= b; i++ {
		// 	hasil = hasil*a
		// }
		for b > 0 {
			if b%2 == 1 {
				hasil *= a
				fmt.Println(hasil)
			}
			a *= a
			fmt.Println(hasil, a)
			b/=2
		}
	}else {
		// for i := b; i<0; i++ {
		// 	hasil = hasil/a
		// 	fmt.Printf("%s %.16f\n","hasil : ", hasil)
		// }
		for b < 0 {
			// fmt.Println("b :",b)
			if b%2 == -1{
				hasil /= a
			}
			// else {
			// 	hasil /= a
			// 	fmt.Println("hasil2 :", hasil)
			// }
			a *= a
			b/=2
		}
	}
	return hasil
}

func main(){
	var bilangan float64
	var pangkat int

	fmt.Printf("Masukkan nilai bilangan : ")
	fmt.Scanln(&bilangan)
	
	fmt.Printf("Masukkan nilai pangkat : ")
	fmt.Scanln(&pangkat)
	
	fmt.Println(pangkat1(bilangan, pangkat))
}

// import "strconv"

// func main(){
// 	var x, n int
// 	var input string
// 	var err error

// 	fmt.Printf("masukkan nilai x : ")
// 	fmt.Scanln(&input)
// 	x, err = strconv.Atoi(input)
// 	if err != nil {
// 		fmt.Println("nilai harus angka")
// 		return
// 	}
	
// 	fmt.Printf("masukkan nilai n : ")
// 	fmt.Scanln(&input)
// 	n, err = strconv.Atoi(input)
// 	if err != nil {
// 		fmt.Println("nilai harus angka")
// 		return
// 	}

// 	fmt.Println()
// }