// // You can edit this code!
// // Click here and start typing.
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	c1 := make(chan int)

// 	go func(c chan int) {
// 		fmt.Println("starts sending data into the channel")
// 		c <- 10
// 		fmt.Println("end sending data into the channel")
// 	}(c1)

// 	fmt.Println("Sleeps for 2 seconds")
// 	time.Sleep(2 * time.Second)

// 	fmt.Println("main go routine starts receiving data")
// 	d := <-c1
// 	fmt.Println("main go routine data:", d)

// 	close(c1)
// 	time.Sleep(time.Second)

// }
