package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 50)

	go func() {
		for i := 50; i >= 0; i-- {
			ch <- i
		}
		close(ch)
	}()

	time.Sleep(1 * time.Second)

	for data := range ch {
		fmt.Println(data)
	}

	fmt.Println("selesai")
}
