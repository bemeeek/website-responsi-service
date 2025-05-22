package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"website-responsi/serial" // Pastikan path ini sesuai dengan struktur proyek Anda

	"github.com/rs/cors" // CORS middleware
)

func main() {
	// Inisialisasi serial port
	if err := serial.InitSerial(); err != nil {
		log.Fatalf("Failed to initialize serial port: %v", err)
	}

	// Membuat mux untuk menangani route
	mux := http.NewServeMux()

	// Endpoint untuk mendapatkan data dari serial
	mux.HandleFunc("/responsi", func(w http.ResponseWriter, r *http.Request) {
		// Membaca data dari port serial
		data, err := serial.ReadData()
		if err != nil {
			// Log error yang lebih mendetail
			log.Printf("Error reading data from serial: %v", err)
			http.Error(w, fmt.Sprintf("Failed to read data from serial: %v", err), http.StatusInternalServerError)
			return
		}

		// Log data yang diterima untuk debugging
		log.Printf("Data received: %v", data)

		// Mengirimkan data sebagai response JSON
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(data)
		if err != nil {
			// Log error encoding JSON
			log.Printf("Error encoding data to JSON: %v", err)
			http.Error(w, fmt.Sprintf("Failed to encode data to JSON: %v", err), http.StatusInternalServerError)
		}
	})

	// Menambahkan CORS middleware untuk mengizinkan akses dari frontend
	handler := cors.Default().Handler(mux)

	// Menjalankan server HTTP pada port 3000
	server := &http.Server{
		Addr:         ":3000",
		Handler:      handler,
		ReadTimeout:  10 * time.Second, // Timeout pembacaan request
		WriteTimeout: 10 * time.Second, // Timeout penulisan response
	}

	log.Println("Server berjalan pada http://localhost:3000")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Server gagal dijalankan:", err)
	}
}
