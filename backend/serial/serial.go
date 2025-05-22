package serial

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/tarm/serial"
)

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

var (
	port    *serial.Port
	scanner *bufio.Scanner
)

// InitSerial membuka port serial dan mempersiapkan scanner untuk pembacaan data
func InitSerial() error {
	config := &serial.Config{
		Name:        "COM7",          // Gantilah dengan port yang sesuai
		Baud:        500000,          // Sesuaikan dengan baud rate perangkat Anda
		ReadTimeout: time.Second * 5, // Timeout pembacaan serial (5 detik)
	}

	var err error
	port, err = serial.OpenPort(config)
	if err != nil {
		log.Printf("Error opening serial port: %v", err) // Log jika ada masalah membuka port
		return fmt.Errorf("failed to open serial port: %w", err)
	}

	// Inisialisasi scanner untuk membaca data dari port serial
	scanner = bufio.NewScanner(port)
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		if i := strings.Index(string(data), "\n"); i >= 0 {
			return i + 1, data[0:i], nil
		}
		return 0, nil, nil
	})

	log.Println("Serial port initialized successfully.")
	return nil
}

// CloseSerial menutup port serial
func CloseSerial() {
	if port != nil {
		port.Close()
		log.Println("Serial port closed.")
	}
}

// ReadData membaca data dari serial port dan memparsingnya ke dalam struct PushButtonData
func ReadData() (PushButtonData, error) {
	data := PushButtonData{}

	// Membaca data dari scanner
	if !scanner.Scan() {
		log.Printf("Error reading from scanner: %v", scanner.Err()) // Log error pembacaan
		return data, fmt.Errorf("scanner failed: %w", scanner.Err())
	}

	strData := strings.TrimSpace(scanner.Text())
	log.Printf("Raw data received: %s", strData) // Debug logging untuk melihat data yang diterima

	// Memecah data berdasarkan koma
	parts := strings.Split(strData, ",")
	log.Printf("Parsed data parts: %v", parts) // Log data yang diparsing

	// Jika ada lebih dari 18 bagian, kita hanya ambil 18 bagian pertama
	if len(parts) > 18 {
		log.Printf("More than 18 parts, trimming extra: %v", parts[0:18])
		parts = parts[:18]
	}

	// Jika ada kurang dari 18 bagian, kita kembalikan data default
	if len(parts) != 18 {
		log.Printf("Invalid data format, expected 18 parts, got %d: %v", len(parts), parts) // Log format yang salah
		return data, fmt.Errorf("invalid data format, expected 18 parts, got %d: %v", len(parts), parts)
	}

	// Fungsi untuk parsing string ke integer
	parseField := func(s string) (int, error) {
		val, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			log.Printf("Invalid number format: %v", err) // Log jika ada kesalahan format
			return 0, fmt.Errorf("invalid number format: %w", err)
		}
		return val, nil
	}

	// Parsing nilai X (button states: 0, 1, 2, 3)
	for i := 0; i < 9; i++ {
		val, err := parseField(parts[i*2])
		if err != nil {
			log.Printf("Invalid X%d value: %v", i+1, err) // Log jika nilai X tidak valid
			return data, fmt.Errorf("invalid X%d value: %w", i+1, err)
		}
		if val < 0 || val > 3 {
			log.Printf("X%d value out of range (0-3): %d", i+1, val) // Log jika nilai X tidak valid
			return data, fmt.Errorf("x%d value out of range (0-3): %d", i+1, val)
		}

		switch i {
		case 0:
			data.X1 = val
		case 1:
			data.X2 = val
		case 2:
			data.X3 = val
		case 3:
			data.X4 = val
		case 4:
			data.X5 = val
		case 5:
			data.X6 = val
		case 6:
			data.X7 = val
		case 7:
			data.X8 = val
		case 8:
			data.X9 = val
		}
	}

	// Parsing nilai Y (waktu dalam ms)
	for i := 0; i < 9; i++ {
		val, err := parseField(parts[i*2+1])
		if err != nil {
			log.Printf("Invalid Y%d value: %v", i+1, err) // Log jika nilai Y tidak valid
			return data, fmt.Errorf("invalid Y%d value: %w", i+1, err)
		}
		if val < 0 {
			log.Printf("Y%d value cannot be negative: %d", i+1, val) // Log jika nilai Y negatif
			return data, fmt.Errorf("y%d value cannot be negative: %d", i+1, val)
		}
		switch i {
		case 0:
			data.Y1 = val
		case 1:
			data.Y2 = val
		case 2:
			data.Y3 = val
		case 3:
			data.Y4 = val
		case 4:
			data.Y5 = val
		case 5:
			data.Y6 = val
		case 6:
			data.Y7 = val
		case 7:
			data.Y8 = val
		case 8:
			data.Y9 = val
		}
	}

	return data, nil
}
