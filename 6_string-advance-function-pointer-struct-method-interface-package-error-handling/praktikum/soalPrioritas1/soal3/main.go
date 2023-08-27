package main

import (
	"fmt"
)

func Compare(a, b string) string {
	var short, long, res string
	if len(a) < len(b) {
		short = a
		long = b
	} else {
		short = b
		long = a
	}

	for i := 0; i < len(long); i++ {
		if i < len(short) && short[i] == long[i] {
			res += string(long[i])
		}
	}
	return res
}

func main() {

	fmt.Println(Compare("AKA", "AKASHI")) //AKA

	fmt.Println(Compare("KANGOORO", "KANG")) //KANG

	fmt.Println(Compare("KI", "KIJANG")) //KI

	fmt.Println(Compare("KUPU-KUPU", "KUPU")) //KUPU

	fmt.Println(Compare("ILALANG", "ILA")) //ILA

}
