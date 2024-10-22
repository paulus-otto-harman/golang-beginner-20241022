package main

import (
	"fmt"
	"time"
)

func suhu(ch chan string) {
	i := 0
	for {
		i++
		time.Sleep(1 * time.Second)
		ch <- fmt.Sprintf("data suhu %d\n", i)
	}
}

func kelembaban(ch chan string) {
	i := 0
	for {
		i++
		time.Sleep(2 * time.Second)
		ch <- fmt.Sprintf("data kelembaban %d\n", i)
	}
}
func tekanan(ch chan string) {
	i := 0
	for {
		i++
		time.Sleep(3 * time.Second)
		ch <- fmt.Sprintf("data tekanan %d\n", i)
	}

}

func main() {
	s := make(chan string, 3)
	k := make(chan string, 3)
	p := make(chan string, 3)

	go suhu(s)
	go kelembaban(k)
	go tekanan(p)

	ticker := time.NewTicker(2 * time.Second)
	done := make(chan bool)
	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()

	go func() {
		for {
			select {
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
				fmt.Println(<-s, <-k, <-p)
			case <-done:
				ticker.Stop()
				return
			}
		}
	}()

	time.Sleep(30 * time.Second)

}
