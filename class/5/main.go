package main

import (
	"20241022/class/5/models"
	"fmt"
	"time"
)

func sensor(nama string, warna int, ch chan models.Sensor, interval time.Duration) {
	satelit := models.InitSensor(f(Color, nama, warna))
	i := 0
	for {
		i++
		time.Sleep(interval * time.Second)
		satelit.Data(i)
		ch <- satelit
	}
}

func main() {
	s := make(chan models.Sensor, 3)
	k := make(chan models.Sensor, 3)
	p := make(chan models.Sensor, 3)

	go sensor("suhu", Green, s, 10)
	go sensor("kelembaban", Blue, k, 2)
	go sensor("tekanan", Yellow, p, 3)

	ticker := time.NewTicker(2 * time.Second)
	done := make(chan bool)
	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()

	go pembacaSensor(s, ticker, done)
	go pembacaSensor(k, ticker, done)
	go pembacaSensor(p, ticker, done)

	time.Sleep(20 * time.Second)

}

func pembacaSensor(ch chan models.Sensor, ticker *time.Ticker, done chan bool) {
	for {
		select {
		case t := <-ticker.C:
			satelit := <-ch
			fmt.Printf("detik ke %-55s pembacaan %s\n", t.String(), satelit.BacaData())

		//	responTerakhir = time.Now().Add(5 * time.Second)
		//case <-time.After(responTerakhir):

		case <-time.After(5 * time.Second):
			satelit := <-ch
			fmt.Printf("detik ke %-55s sensor %s : %s\n", time.Now().String(), f(Bold, satelit.NamaSensor()), f(Color, "mengalami timeout 5 detik", Red))
		case <-done:
			ticker.Stop()
			return
		}
	}
}

const Red = 31
const Green = 92
const Yellow = 93
const Blue = 94

const Bold = "\033[1m%s\033[0m"
const Color = "\x1b[%dm%s\x1b[0m"

func f(mode string, teks string, warna ...int) string {

	switch {
	case mode == Bold && len(warna) > 0:
		return fmt.Sprintf(Color, warna[0], fmt.Sprintf(Bold, teks))
	case mode == Bold:
		return fmt.Sprintf(Bold, teks)
	default:
		return fmt.Sprintf(Color, warna[0], teks)
	}
}
