package main

import "fmt"

func roman(num int) string {
	var converted string
	for num > 0 {
		if num == 0 {
			return converted
		}
		if num >= 1000 {
			converted += "M"
			num -= 1000
		}
		if num >= 900 && num < 1000 {
			converted += "CM"
			num -= 900
		}
		if num >= 500 && num < 900 {
			converted += "D"
			num -= 500
		}
		if num >= 400 && num < 500 {
			converted += "CD"
			num -= 400
		}
		if num >= 100 && num < 400 {
			converted += "C"
			num -= 100
		}
		if num >= 90 && num < 100 {
			converted += "XL"
			num -= 90
		}
		if num >= 50 && num < 90 {
			converted += "L"
			num -= 50
		}
		if num >= 40 && num < 50 {
			converted += "XL"
			num -= 40
		}
		if num >= 10 && num < 40 {
			converted += "X"
			num -= 10
		}
		if num == 9 {
			converted += "IX"
			num -= 9
		}
		if num >= 5 && num < 9 {
			converted += "V"
			num -= 5
		}
		if num == 4 {
			converted += "IV"
			num -= 4
		}
		if num >= 1 && num < 4 {
			converted += "I"
			num -= 1
		}
	}
	return converted
}

func main() {
	fmt.Println(roman(4))
	fmt.Println(roman(9))
	fmt.Println(roman(23))
	fmt.Println(roman(2021))
	fmt.Println(roman(1646))
}
