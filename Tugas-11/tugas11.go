package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

// Soal Nomor 2
func getMovies(ch chan string, movies ...string) {
	for _, movie := range movies {
		ch <- movie // Mengirim data movie ke channel
	}
	close(ch) // Menutup channel setelah semua data dikirim
}

// Soal Nomor 3
type CalculationResult struct {
	Radius        float64
	Area          float64
	Circumference float64
	Volume        float64
}

func calculate(radius float64, height float64, ch chan CalculationResult) {
	area := math.Pi * radius * radius
	circumference := 2 * math.Pi * radius
	volume := area * height

	result := CalculationResult{
		Radius:        radius,
		Area:          area,
		Circumference: circumference,
		Volume:        volume,
	}

	ch <- result
}

// Soal Nomor 4
type GeometryResult struct {
	Length    float64
	Width     float64
	Height    float64
	Area      float64
	Perimeter float64
	Volume    float64
}

func computeGeometry(length, width, height float64, resultChan chan GeometryResult) {
	result := GeometryResult{
		Length:    length,
		Width:     width,
		Height:    height,
		Area:      length * width,
		Perimeter: 2 * (length + width),
		Volume:    length * width * height,
	}
	resultChan <- result
}

func main() {
	// Jawaban nomor 1
	phones := []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}

	var wg sync.WaitGroup

	for idx, phone := range phones {
		wg.Add(1)

		go func(idx int, phone string) {
			defer wg.Done()
			time.Sleep(time.Duration(idx) * time.Second)

			fmt.Printf("%d. %s\n", idx+1, phone)
		}(idx, phone)
	}

	wg.Wait()

	// Jawaban Nomor 2
	movies := []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}

	moviesChannel := make(chan string)

	go getMovies(moviesChannel, movies...)

	for value := range moviesChannel {
		fmt.Println(value)
	}

	// Jawaban Nomor 3
	radii := []float64{8, 14, 20}
	height := 10.0

	ch := make(chan CalculationResult)

	for _, radius := range radii {
		go calculate(radius, height, ch)
	}

	for range radii {
		result := <-ch
		fmt.Printf("Jari-jari: %.2f\n", result.Radius)
		fmt.Printf("Luas lingkaran: %.2f\n", result.Area)
		fmt.Printf("Keliling lingkaran: %.2f\n", result.Circumference)
		fmt.Printf("Volume tabung: %.2f\n\n", result.Volume)
	}

	close(ch)

	// JAwaban Nomor 4
	length, width, height := 5.0, 3.0, 4.0

	resultChan := make(chan GeometryResult)

	go computeGeometry(length, width, height, resultChan)

	select {
	case result := <-resultChan:
		fmt.Printf("Panjang: %.2f\n", result.Length)
		fmt.Printf("Lebar: %.2f\n", result.Width)
		fmt.Printf("Tinggi: %.2f\n", result.Height)
		fmt.Printf("Luas persegi panjang: %.2f\n", result.Area)
		fmt.Printf("Keliling persegi panjang: %.2f\n", result.Perimeter)
		fmt.Printf("Volume balok: %.2f\n", result.Volume)
	}

}
