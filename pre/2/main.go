package main

import "fmt"

func main() {
	ch := make(chan int)

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("test")
		}()
	}

	fmt.Println(ch)
}
