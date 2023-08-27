package main

import (
	"fmt"
	// "praktikum/praktikum/soalPrioritas2/soal1/model"
)

func caesar(offset int, input string) string {
	var ascii string
	for _, character := range input {
		modOff := offset % 26
		tmp := int(character) + modOff
		if tmp > 122 {
			tmp -= 26
		}
		ascii += fmt.Sprintf("%c", tmp)
	}
	return ascii
}

func main() {

	fmt.Println(caesar(3, "abc")) // def

	fmt.Println(caesar(2, "alta")) // def

	fmt.Println(caesar(10, "alterraacademy")) // kvdobbkkmknowi

	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz")) // bcdefghijklmnopqrstuvwxyza

	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl

}

// func main() {
// 	model := new(model.ASCI)

// 	fmt.Println(model.ConvertASCII("abc", 3))
// 	fmt.Println(model.ConvertASCII("abcdefghijklmnopqrstuvwxyz", 1))
// 	fmt.Println(model.ConvertASCII("abcdefghijklmnopqrstuvwxyz", 1000))
// }
