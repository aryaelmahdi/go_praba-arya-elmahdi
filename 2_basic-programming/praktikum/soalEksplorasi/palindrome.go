package main

import "fmt"

func main(){
	var arrange, input string

	fmt.Printf("Masukkan Kata : ")
	fmt.Scanln(&input)

	for i:= (len(input)-1) ; i>=0 ; i-- {
		if i == (len(input)-1) {
			arrange = string(input[i])
		}else {
			arrange = arrange + string(input[i])
		}
	}

	fmt.Println(input, ":", arrange)
	if arrange == input {
		fmt.Println(input, " adalah Palindrome")
	}else {
		fmt.Println(input, " bukan Palindrome")
	}
}