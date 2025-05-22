package main

import (
	"fmt"
	"log"

	"github.com/tarm/serial"
)

func main() {
	// Konfigurasi port serial
	config := &serial.Config{
		Name:     "COM7", // Gantilah dengan port serial yang sesuai, misalnya COM1, COM3
		Baud:     500000, // Sesuaikan baud rate dengan Arduino atau perangkat Anda
		Size:     8,      // Ukuran data 8 bit
		Parity:   serial.ParityNone,
		StopBits: 1, // 1 stop bit
	}

	// Membuka port serial
	port, err := serial.OpenPort(config)
	if err != nil {
		log.Fatalf("Failed to open serial port: %v", err)
		return
	}
	defer port.Close()

	// Membaca data dari port serial
	buf := make([]byte, 128) // Buffer untuk membaca data
	for {
		n, err := port.Read(buf)
		if err != nil {
			log.Printf("Error reading from serial port: %v", err)
			break
		}

		if n > 0 {
			// Menampilkan data yang dibaca dari serial
			fmt.Printf("Data read: %s\n", string(buf[:n]))
		}
	}
}
