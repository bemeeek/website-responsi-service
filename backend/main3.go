package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/tarm/serial"
)

var port *serial.Port
var mqttClient mqtt.Client
var scanner *bufio.Scanner

// Struktur data tombol yang akan diparse
type PushButtonData struct {
	X1 int `json:"x1"` // Nomor tombol yang ditekan pertama (urutan x=1)
	Y1 int `json:"y1"`
	X2 int `json:"x2"` // Nomor tombol yang ditekan kedua (urutan x=2)
	Y2 int `json:"y2"`
	X3 int `json:"x3"` // Nomor tombol yang ditekan ketiga (urutan x=3)
	Y3 int `json:"y3"`
}

// Fungsi untuk inisialisasi serial
func InitSerial() error {
	config := &serial.Config{
		Name:        "COM7",          // Gantilah dengan port yang sesuai
		Baud:        500000,          // Sesuaikan dengan baud rate perangkat Anda
		ReadTimeout: time.Second * 1, // Timeout pembacaan serial (1 detik)
	}

	var err error
	port, err = serial.OpenPort(config)
	if err != nil {
		log.Printf("Error opening serial port: %v", err)
		return fmt.Errorf("failed to open serial port: %w", err)
	}
	log.Println("Serial port initialized successfully.")

	// Inisialisasi scanner untuk membaca data
	scanner = bufio.NewScanner(port)
	return nil
}

// Fungsi untuk menginisialisasi MQTT client
func InitMQTT() mqtt.Client {
	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:8888").SetClientID("arduino_client")
	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Failed to connect to MQTT broker: %v", token.Error())
	}
	return client
}

// Fungsi untuk membaca data dari serial port sampai tanda pagar
func ReadData() (PushButtonData, error) {
	data := PushButtonData{}

	for scanner.Scan() {
		strData := scanner.Text()
		log.Printf("Raw data: %s", strData)

		endIndex := strings.Index(strData, "#")
		if endIndex != -1 {
			dataUntilHash := strData[:endIndex]
			log.Printf("Data until # : %s", dataUntilHash)

			parts := strings.Split(dataUntilHash, ";")
			for i, part := range parts {
				part = strings.TrimSpace(part)
				if part == "" || part == "[0,0]" {
					continue // Lewati data kosong
				}

				// Hilangkan karakter '[' dan ']'
				part = strings.TrimPrefix(part, "[")
				part = strings.TrimSuffix(part, "]")

				// Pisahkan x dan y
				xy := strings.Split(part, ",")
				if len(xy) != 2 {
					continue // Format tidak valid
				}

				xStr := strings.TrimSpace(xy[0])
				yStr := strings.TrimSpace(xy[1])

				x, errX := strconv.Atoi(xStr)
				y, errY := strconv.Atoi(yStr)

				if errX != nil || errY != nil || x < 1 {
					continue // Data tidak valid
				}

				// Nomor tombol = posisi dalam data (indeks + 1)
				buttonNumber := i + 1

				// Isi struct sesuai urutan tekan (x)
				switch x {
				case 1:
					data.X1 = buttonNumber
					data.Y1 = y
				case 2:
					data.X2 = buttonNumber
					data.Y2 = y
				case 3:
					data.X3 = buttonNumber
					data.Y3 = y
				}
			}

			return data, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return data, err
	}

	return data, fmt.Errorf("no '#' found in data")
}

// Fungsi untuk mengirimkan data ke topik MQTT dalam format JSON
func PublishDataToMQTT(data PushButtonData) {
	topic := "arduino/data"

	// Mengonversi data ke format JSON
	payload, err := json.Marshal(data)
	if err != nil {
		log.Printf("Failed to marshal data to JSON: %v", err)
		return
	}

	// Mempublikasikan data ke MQTT
	token := mqttClient.Publish(topic, 0, false, payload)
	token.Wait()
	log.Printf("Data sent to MQTT: %s", payload)
}

func main() {
	// Inisialisasi serial port dan MQTT client
	if err := InitSerial(); err != nil {
		log.Fatalf("Failed to initialize serial port: %v", err)
	}
	mqttClient = InitMQTT()
	defer mqttClient.Disconnect(250)
	defer port.Close()

	// Loop untuk membaca dan mengirim data ke MQTT
	for {
		data, err := ReadData()
		if err != nil {
			log.Println("Error reading data:", err)
			continue
		}

		// Mengirimkan data ke MQTT
		PublishDataToMQTT(data)

		// time.Sleep(0.1) // Delay 1 detik sebelum membaca lagi
	}
}
