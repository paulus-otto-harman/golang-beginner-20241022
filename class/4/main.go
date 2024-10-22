package main

import (
	"20241022/class/4/models"
	"fmt"
	"sync"
)

var mutex = &sync.Mutex{}

func kurangiStok(wg *sync.WaitGroup, produk *models.Product, jumlah int) {
	defer wg.Done()

	mutex.Lock()
	if kuantity := produk.Quantity(); kuantity > 0 {
		fmt.Printf("sebelum dikurangi : %d\n", kuantity)
		produk.SetQuantity(kuantity - jumlah)
		fmt.Printf("setelah dikurangi : %d\n", produk.Quantity())

	} else {
		fmt.Println("stok tidak mencukupi")
	}
	mutex.Unlock()
}

func main() {
	produk := models.InitProduct("laptop", 10)

	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go kurangiStok(&wg, &produk, 1)
	}
	wg.Wait()

	println(produk.Quantity())
}
