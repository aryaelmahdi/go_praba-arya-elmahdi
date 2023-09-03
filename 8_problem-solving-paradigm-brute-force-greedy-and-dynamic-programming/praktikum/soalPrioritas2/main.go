package main

import "fmt"

func Frog(jumps []int) int {
	var cost int
	var current int

	for i := 0; i < len(jumps); i++ {
		current = jumps[i]
		for j := i + 1; j < len(jumps); j++ {
			if len(jumps) > 4 {
				if i+1 < len(jumps) && jumps[i]-jumps[j+1] <= jumps[i]-jumps[j] {
					cost -= current - jumps[i+2]
					i += 2
				}
				if i+1 < len(jumps) && jumps[i]-jumps[j] > jumps[i]-jumps[j+1] {
					cost -= jumps[i] - jumps[i+2]
					i += 2
				}
				if i == len(jumps)-2 {
					cost += jumps[i] - jumps[i+1]
					i++
				}
			}
			if j+1 < len(jumps) && len(jumps) <= 4 {
				if i+1 < len(jumps) && jumps[i]-jumps[j] > jumps[i]-jumps[j+1] {
					cost -= current + jumps[i+1]
					i++
				}
			}

			if i+2 < len(jumps) {
				cost -= current - jumps[i+2]
				i += 2
			}

			if i == len(jumps)-2 {
				cost -= current - jumps[i+1]
				i++
			}
		}
	}
	if cost < 0 {
		return -cost
	}
	return cost
}

func main() {

	fmt.Println(Frog([]int{10, 30, 40, 20})) // 30

	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40

}
