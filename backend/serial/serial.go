package serial

import (
	"log"
	"time"

	"github.com/tarm/serial"
)

// Struktur untuk data yang diterima dari Arduino
type PushButtonData struct {
	X1 int `json:"x1"`
	X2 int `json:"x2"`
	X3 int `json:"x3"`
	X4 int `json:"x4"`
	X5 int `json:"x5"`
	X6 int `json:"x6"`
	X7 int `json:"x7"`
	X8 int `json:"x8"`
	X9 int `json:"x9"`
	Y1 int `json:"y1"`
	Y2 int `json:"y2"`
	Y3 int `json:"y3"`
	Y4 int `json:"y4"`
	Y5 int `json:"y5"`
	Y6 int `json:"y6"`
	Y7 int `json:"y7"`
	Y8 int `json:"y8"`
	Y9 int `json:"y9"`
}

// Fungsi untuk membaca data dari Arduino melalui port serial
func ReadData() (PushButtonData, error) {
	// Konfigurasi port serial
	config := &serial.Config{
		Name:     "COM3", // Gantilah dengan port serial yang sesuai
		Baud:     9600,
		Size:     8,
		Parity:   serial.ParityNone, // Pastikan menggunakan serial.ParityNone
		StopBits: 1,                 // Pastikan menggunakan 1 stop bit
	}

	// Membuka port serial
	port, err := serial.OpenPort(config)
	if err != nil {
		log.Fatal(err)
		return PushButtonData{}, err
	}

	// Membaca data dari port serial
	buf := make([]byte, 128)
	_, err = port.Read(buf)
	if err != nil {
		log.Fatal(err)
		return PushButtonData{}, err
	}

	// Misalkan kita menerima data berupa "X1,X2,X3,...,Y1,Y2,Y3,..."
	// Parsing data tersebut ke dalam struktur
	data := PushButtonData{
		X1: int(buf[0] - '0'),
		X2: int(buf[1] - '0'),
		X3: int(buf[2] - '0'),
		X4: int(buf[3] - '0'),
		X5: int(buf[4] - '0'),
		X6: int(buf[5] - '0'),
		X7: int(buf[6] - '0'),
		X8: int(buf[7] - '0'),
		X9: int(buf[8] - '0'),
		Y1: int(time.Now().Unix()),
		Y2: int(time.Now().Unix()),
		Y3: int(time.Now().Unix()),
		Y4: int(time.Now().Unix()),
		Y5: int(time.Now().Unix()),
		Y6: int(time.Now().Unix()),
		Y7: int(time.Now().Unix()),
		Y8: int(time.Now().Unix()),
		Y9: int(time.Now().Unix()),
	}

	return data, nil
}
