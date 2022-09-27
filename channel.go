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

	//Select
	queue := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			queue <- i
		}

		done <- true
	}()

	for {
		select {
		case v := <-queue:
			fmt.Println(v)
		case <-done:
			fmt.Println("done.")
			return
		}
	}
}
