package main

import "fmt"

// Unbuffered  channel

func main() {
	//Unbuffered  channel
	ch := make(chan int)

	go func() {
		ch <- 100
	}()

	fmt.Println(<-ch)

	fmt.Println("done.")
}
