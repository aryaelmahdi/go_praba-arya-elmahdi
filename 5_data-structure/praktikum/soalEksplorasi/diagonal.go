package main

import "fmt"

func absoluteSum(arr [][]int, target int) int{
	var sumA, sumB int = 0, 0
	for i:= 0; i<target; i++ {
		for j:= 0; j<target; j++{
			if i == j {
				sumA = sumA + arr[i][j]
			}
		}
	}
	for i:= (target-1); i>=0; i--{
		for j:= (target-1); j>=0; j--{
			if i + j == target-1 {
				sumB = sumB + arr[i][j]
			}
		}
	}
	if sumA - sumB < 0 {
		return -(sumA - sumB)
	}
	return sumA - sumB
}

func main(){
	target := 3
	mat := [][]int {
		{1,2,3},
		{4,5,6},
		{9,8,9},
	}
	fmt.Println(absoluteSum(mat, target))
}