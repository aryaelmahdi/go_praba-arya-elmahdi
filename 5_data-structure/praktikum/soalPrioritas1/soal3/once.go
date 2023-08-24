package main


import (
"fmt"
"strconv"
)


func munculSekali(angka string) []int {
	var hasil []int
	var err error
	var temp int
	for i:= 0; i<len(angka); i++{
		found := false
		for j:=0; j<len(angka); j++{
			if j==i {
				continue
			}else {
				if angka[j] == angka[i] {
					found = true
				}
			}
		}
		if !found {
			temp, err = strconv.Atoi(string(angka[i]))
			if err == nil {
				hasil = append(hasil, temp)
			}
		}
	}
	return hasil
}


func main() {

// Test cases

fmt.Println(munculSekali("1234123")) // [4]

fmt.Println(munculSekali("76523752")) // [6 3]

fmt.Println(munculSekali("12345")) // [1 2 3 4 5]

fmt.Println(munculSekali("1122334455")) // []

fmt.Println(munculSekali("0872504")) // [8 7 2 5 4]

}