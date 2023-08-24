package main

import "fmt"

func ArrayMerge(a []string, b []string) []string{
	for i:=0; i<len(b); i++ {
		found := false
		for j:= 0; j<len(a); j++ {
			if a[j] == b[i] {
				found = true
			}
		}
		if !found {
			a = append(a, b[i])
		}
	}
	return a
}

func main() {

	// Test cases
	
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	
	// ["king", "devil jin", "akuma", "eddie", "steve", "geese"]
	
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	
	// ["sergei", "jin", "steve", "bryan"]
	
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	
	// ["alisa", "yoshimitsu", "devil jin", "law"]
	
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	
	// ["devil jin", "sergei"]
	
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	
	// ["hwoarang"]
	
	fmt.Println(ArrayMerge([]string{}, []string{}))
	
	// []
	
	}