package main

import "fmt"

func main() {
	//Unbuffered  channel

	// ch := make(chan int)

	// go func() {
	// 	ch <- 100
	// }()

	// fmt.Println(<-ch)

	// fmt.Println("done.")

	//Buffered channel
	ch := make(chan int, 2)

	ch <- 1
	ch <- 2
	close(ch)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
