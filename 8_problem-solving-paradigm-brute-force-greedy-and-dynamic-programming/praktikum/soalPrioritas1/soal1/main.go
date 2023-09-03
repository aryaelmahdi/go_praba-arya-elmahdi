package main

import (
	"fmt"
	"strconv"
)

var (
	res []string
)

func toBinary(num int) []string {
	if len(res) > num {
		return res[:num+1]
	}

	for i := len(res); i <= num; i++ {
		binary := strconv.FormatInt(int64(i), 2)
		res = append(res, binary)
	}
	return res
}

func main() {
	fmt.Println(toBinary(2))
	fmt.Println(toBinary(3))
}
