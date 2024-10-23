package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	go func() {
		ch <- 1
		close(ch)
	}()
	go func() {
		ch <- 1
		close(ch)
	}()
	fmt.Println(<-ch)
}
