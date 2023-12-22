package main

import (
	"flag"
	"fmt"
	"math"
	"sort"
	"time"
)

// Soal 1
func printInfo(kalimat string, tahun int) {
	defer fmt.Printf("Kalimat: %s, Tahun: %d\n", kalimat, tahun)
}

// Soal 2
func kelilingSegitigaSamaSisi(sisi int, display bool) (string, error) {
	if sisi == 0 {
		if display {
			return "", fmt.Errorf("Maaf anda belum menginput sisi dari segitiga sama sisi")
		}
		return "", fmt.Errorf("Maaf anda belum menginput sisi dari segitiga sama sisi")
	}

	keliling := 3 * sisi
	if display {
		return fmt.Sprintf("keliling segitiga sama sisinya dengan sisi %d cm adalah %d cm", sisi, keliling), nil
	}
	return fmt.Sprintf("%d", keliling), nil
}

// Soal 3
func tambahAngka(n int, angka *int) {
	*angka += n
}

func cetakAngka(angka *int) {
	fmt.Println("Total angka:", *angka)
}

// Soal 4
func populatePhones(phones *[]string) {
	*phones = []string{"Xiaomi", "Asus", "IPhone", "Samsung", "Oppo", "Realme", "Vivo"}
	sort.Strings(*phones)

	for i, phone := range *phones {
		fmt.Printf("%d. %s\n", i+1, phone)
		time.Sleep(1 * time.Second)
	}
}

// Soal 5
func luasLingkaran(jariJari float64) int {
	return int(math.Round(math.Pi * math.Pow(jariJari, 2)))
}

func kelilingLingkaran(jariJari float64) int {
	return int(math.Round(2 * math.Pi * jariJari))
}

// Soal 6
func main() {
	// Jawaban nomor 1
	printInfo("Golang Backend Development", 2021)

	// Jawaban nomor 2
	fmt.Println(kelilingSegitigaSamaSisi(4, true))
	fmt.Println(kelilingSegitigaSamaSisi(8, false))
	fmt.Println(kelilingSegitigaSamaSisi(0, true))
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic recovered:", r)
		}
	}()
	fmt.Println(kelilingSegitigaSamaSisi(0, false))

	// Jawaban nomor 3
	angka := 1
	defer cetakAngka(&angka)
	tambahAngka(7, &angka)
	tambahAngka(6, &angka)
	tambahAngka(-1, &angka)
	tambahAngka(9, &angka)

	// Jawaban nomor 4
	var phones []string
	populatePhones(&phones)

	// Jawaban nomor 5
	fmt.Println("Luas lingkaran dengan jari-jari 7:", luasLingkaran(7))
	fmt.Println("Keliling lingkaran dengan jari-jari 7:", kelilingLingkaran(7))
	fmt.Println("Luas lingkaran dengan jari-jari 10:", luasLingkaran(10))
	fmt.Println("Keliling lingkaran dengan jari-jari 10:", kelilingLingkaran(10))
	fmt.Println("Luas lingkaran dengan jari-jari 15:", luasLingkaran(15))
	fmt.Println("Keliling lingkaran dengan jari-jari 15:", kelilingLingkaran(15))

	// Jawaban nomor 6
	var panjang, lebar int
	flag.IntVar(&panjang, "panjang", 0, "Panjang persegi panjang")
	flag.IntVar(&lebar, "lebar", 0, "Lebar persegi panjang")
	flag.Parse()

	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)
	fmt.Println("Luas persegi panjang:", luas)
	fmt.Println("Keliling persegi panjang:", keliling)
}
