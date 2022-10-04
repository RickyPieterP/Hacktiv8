package main

import (
	"fmt"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)
	c4 := make(chan int)
	defer close(c1)
	defer close(c2)
	defer close(c3)
	defer close(c4)

	go intro(10, c1)
	go intro(11, c2)
	go intro(12, c3)
	go intro(13, c4)

	result1 := <-c1
	fmt.Println(result1)
	
	result2 := <-c2
	fmt.Println(result2)

	result3 := <-c3
	fmt.Println(result3)

	result4 := <-c4
	fmt.Println(result4)
}

func intro(a int, c chan int) {
	data := a

	c <- data
}
