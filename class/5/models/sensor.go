package models

import (
	"fmt"
	"time"
)

type Sensor struct {
	nama string
	data string
}

func InitSensor(nama string) Sensor {
	return Sensor{nama: nama}
}

func (s *Sensor) Nama(nama string) {
	s.nama = nama
}

func (s *Sensor) Data(data int) {
	s.data = fmt.Sprintf("data %s %d", s.nama, data)
	fmt.Printf("detik ke %-55s mengirim data %s %d\n", time.Now().String(), s.nama, data)
}

func (s *Sensor) NamaSensor() string {
	return s.nama
}

func (s *Sensor) BacaData() string {
	return s.data
}
