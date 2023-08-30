package main

import (
	"fmt"
)

type pair struct {
	name  string
	count int
}

func MostAppearItem(items []string) []pair {
	resMap := make(map[string]int)
	for _, item := range items {
		if _, exists := resMap[item]; exists {
			resMap[item]++
		} else {
			resMap[item] = 1
		}
	}
	var list []pair
	for name, counts := range resMap {
		list = append(list, pair{name: name, count: counts})
	}

	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list)-i-1; j++ {
			if list[j+1].count < list[j].count {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
	return list
}

func main() {
	pairs := MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}) // golang->1 ruby->2 js->4
	for _, list := range pairs {
		fmt.Print(list.name, " -> ", list.count, " ")
	}

	fmt.Println()
	pairs = MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}) // C->1 D->2 B->3 A->4
	for _, list := range pairs {
		fmt.Print(list.name, " -> ", list.count, " ")
	}

	fmt.Println()
	pairs = MostAppearItem([]string{"football", "basketball", "tenis"}) // football->1 basketball->1 tenis->1
	for _, list := range pairs {
		fmt.Print(list.name, " -> ", list.count, " ")
	}
}
