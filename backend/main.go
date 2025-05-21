package main

import (
	"website-responsi/serial" // Pastikan path ini sesuai dengan struktur proyek Anda

	"github.com/gin-gonic/gin"
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

func main() {
	// Inisialisasi router Gin
	r := gin.Default()

	// Endpoint untuk mendapatkan data push button
	r.GET("/responsi", func(c *gin.Context) {
		// Membaca data dari Arduino melalui serial
		data, err := serial.ReadData()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, data)
	})

	// Jalankan server di port 8080
	r.Run(":8080")
}
