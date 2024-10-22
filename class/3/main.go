package main

import (
	"fmt"
	"runtime"
	"sync"
)

func ganjil(wg *sync.WaitGroup, angka int) {
	defer wg.Done()
	if angka%2 != 0 {
		fmt.Println(angka)
	}

}

func main() {
	fmt.Printf("Jumlah Core CPU : %d \n", runtime.NumCPU())
	fmt.Printf("GOMAXPROCS : %d \n", runtime.GOMAXPROCS(3))

	wg := sync.WaitGroup{}

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go ganjil(&wg, i)
	}
	wg.Wait()

}
