package main

import (
	"fmt"
)

var (
	res [][]int
)

func pascal(row int) (res [][]int) {
	res = make([][]int, row)
	if row < 0 {
		return nil
	}
	if len(res) > row {
		return res[:row+1]
	}
	for i := 0; i < row; i++ {
		res[i] = make([]int, i+1)
		res[i][i], res[i][0] = 1, 1
		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}
	}
	return res
}

func main() {
	fmt.Println(pascal(5))
}
