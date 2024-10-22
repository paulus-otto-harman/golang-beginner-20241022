package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex

func beliTiket(wg *sync.WaitGroup, sisaTiket *int) {
	defer wg.Done()

	mutex.Lock()
	if *sisaTiket > 0 {
		*sisaTiket--

		fmt.Println("tiket berhasil terjual, sisa", *sisaTiket)
	} else {
		fmt.Println("tiket habis terjual")
	}
	mutex.Unlock()
}

func main() {
	wg := sync.WaitGroup{}

	tiketTersedia := 50

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go beliTiket(&wg, &tiketTersedia)
	}

	wg.Wait()
	fmt.Println("sisa tiket : ", tiketTersedia)
}
