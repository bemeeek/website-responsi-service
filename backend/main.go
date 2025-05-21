package main

import (
	"time"
	"website-responsi/serial" // Pastikan path ini sesuai dengan struktur proyek Anda

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var mockButtonData serial.PushButtonData

func main() {
	// Inisialisasi router Gin
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Pastikan frontend Anda ada di sini
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Endpoint untuk mendapatkan data push button
	r.GET("/responsi", func(c *gin.Context) {
		c.JSON(200, mockButtonData)
		// Membaca data dari Arduino melalui serial
		//data, err := serial.ReadData()
		//if err != nil {
		//	c.JSON(500, gin.H{"error": err.Error()})
		//	return
		//}
		//c.JSON(200, data)
	})

	r.POST("/mock", func(c *gin.Context) {
		var data serial.PushButtonData
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		mockButtonData = data
		c.JSON(200, mockButtonData)
		c.JSON(200, gin.H{"message": "Data mock berhasil diterima"})
	})

	// Jalankan server di port 8080
	r.Run(":3000")
}
