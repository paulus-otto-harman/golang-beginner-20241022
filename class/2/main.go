package main

import (
	"fmt"
	"time"
)

func satu(ch chan string, msg string) {
	time.Sleep(9 * time.Millisecond)
	ch <- msg
}
func dua(ch chan string, msg string) {
	time.Sleep(6 * time.Millisecond)
	ch <- msg
}
func tiga(ch chan string, msg string) {
	time.Sleep(3 * time.Millisecond)
	ch <- msg
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)

	go satu(ch1, "satu")
	go dua(ch2, "dua")
	go tiga(ch3, "tiga")

	time.Sleep(3 * time.Second)
	select {
	case data := <-ch1:
		fmt.Println("Menerima:", data)
	case data := <-ch2:
		fmt.Println("Menerima:", data)
	case data := <-ch3:
		fmt.Println("Menerima:", data)
	default:
		fmt.Println("Tidak ada data yang siap untuk diterima.")
		time.Sleep(20 * time.Millisecond) // Simulasi operasi lain
	}

}
