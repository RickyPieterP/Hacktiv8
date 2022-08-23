package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	// Looping
	fmt.Println("1. Looping")
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "Genap")
		} else {
			fmt.Println(i, "Ganjil")
		}
	}

	// Slice
	fmt.Println("=====================")
	fmt.Println("2. Slice")
	friends := []string{"nama1", "nama2", "nama3", "nama4", "nama5", "nama6", "nama7", "nama8", "nama9", "nama10"}

	for _, v := range friends {
		fmt.Println(v)
	}
}
