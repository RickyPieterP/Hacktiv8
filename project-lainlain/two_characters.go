package main

import (
	"fmt"
)

func main() {
	stringSlice := "beabeefeab"
	a := make(map[int]string)
	list := []string{}
	keys := make(map[string]bool)
	alt := [][]int{}

	for i, v := range stringSlice {
		a[i] = string(v)
	}

	for _, entry := range a {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	for i := 0; i < len(list); i++ {
		
	}


	// a b c     = [a, b] [a, c] [b, c] // 3
	// a b c d   = [a, b] [a, c] [a, d] [b, c] [b, d] [c, d] // 6
	// a b c d e = [a, b] [a, c] [a, d] [a, e] [b, c] [b, d] [b, e] [c, d] [c, e] [d, e] // 10
}
