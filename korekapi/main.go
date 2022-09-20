package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	p1 := make(chan int)
	p2 := make(chan int)
	p3 := make(chan int)
	p4 := make(chan int)

	go put(p1)
	go put(p2)
	go put(p3)
	go put(p4)

	i := 1
	for {
		select {
		case v := <-p1:
			fmt.Printf("korek ada di player 1 pada hit ke %d dan mempunyai nilai  :%d\n", i, v)
			check(v, p1, "player1", i)

		case v := <-p2:
			fmt.Printf("korek ada di player 2 pada hit ke %d dan mempunyai nilai  :%d\n", i, v)
			check(v, p2, "player2", i)

		case v := <-p3:
			fmt.Printf("korek ada di player 3 pada hit ke %d dan mempunyai nilai  :%d\n", i, v)
			check(v, p3, "player 3", i)

		case v := <-p4:
			fmt.Printf("korek ada di player 4 pada hit ke %d dan mempunyai nilai  :%d\n", i, v)
			check(v, p4, "player 4", i)

		case <-time.After(10 * time.Second):
			fmt.Printf("TIMEOUT \n...")
			return
		}
		i++
		time.Sleep(1 * time.Second)
	}
}

func put(c chan int) {
	c <- rand.Intn(100) + 1
}

func check(in int, c chan int, p string, hit int) {
	in = in % 11

	if in == 0 {
		fmt.Printf("%s kalah pada hit ke %d\n", p, hit)
		os.Exit(1)
	} else {
		go put(c)
	}
}
