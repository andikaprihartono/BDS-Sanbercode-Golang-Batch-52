package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Soal nomor 1
	panjangPersegiPanjang := "8"
	lebarPersegiPanjang := "5"
	alasSegitiga := "6"
	tinggiSegitiga := "7"

	panjang, _ := strconv.Atoi(panjangPersegiPanjang)
	lebar, _ := strconv.Atoi(lebarPersegiPanjang)
	alas, _ := strconv.Atoi(alasSegitiga)
	tinggi, _ := strconv.Atoi(tinggiSegitiga)

	luasPersegiPanjang := panjang * lebar
	kelilingPersegiPanjang := 2 * (panjang + lebar)
	luasSegitiga := 0.5 * float64(alas) * float64(tinggi)

	// Jawaban Soal nomor 1
	fmt.Println("Luas Persegi Panjang:", luasPersegiPanjang)
	fmt.Println("Keliling Persegi Panjang:", kelilingPersegiPanjang)
	fmt.Println("Luas Segitiga:", luasSegitiga)

	// Soal Nomor 2
	nilaiJohn := 80
	nilaiDoe := 50

	// Jawaban Soal Nomor 2
	fmt.Print("Nilai John: ")
	if nilaiJohn >= 80 {
		fmt.Println("A")
	} else if nilaiJohn >= 70 && nilaiJohn < 80 {
		fmt.Println("B")
	} else if nilaiJohn >= 60 && nilaiJohn < 70 {
		fmt.Println("C")
	} else if nilaiJohn >= 50 && nilaiJohn < 60 {
		fmt.Println("D")
	} else {
		fmt.Println("E")
	}

	fmt.Print("Nilai Doe: ")
	if nilaiDoe >= 80 {
		fmt.Println("A")
	} else if nilaiDoe >= 70 && nilaiDoe < 80 {
		fmt.Println("B")
	} else if nilaiDoe >= 60 && nilaiDoe < 70 {
		fmt.Println("C")
	} else if nilaiDoe >= 50 && nilaiDoe < 60 {
		fmt.Println("D")
	} else {
		fmt.Println("E")
	}

	// Soal Nomor 3
	tanggal := 17
	bulan := 8
	tahun := 1945

	var namaBulan string
	switch bulan {
	case 1:
		namaBulan = "Januari"
	case 2:
		namaBulan = "Februari"
	case 3:
		namaBulan = "Maret"
	case 4:
		namaBulan = "April"
	case 5:
		namaBulan = "Mei"
	case 6:
		namaBulan = "Juni"
	case 7:
		namaBulan = "Juli"
	case 8:
		namaBulan = "Agustus"
	case 9:
		namaBulan = "September"
	case 10:
		namaBulan = "Oktober"
	case 11:
		namaBulan = "November"
	case 12:
		namaBulan = "Desember"
	default:
		namaBulan = "Bulan tidak valid"
	}

	// Jawaban Soal Nomor 3
	fmt.Printf("%d %s %d\n", tanggal, namaBulan, tahun)

	// Soal Nomor 4
	tahunKelahiran := 1992
	var generasi string
	switch {
	case tahunKelahiran >= 1944 && tahunKelahiran <= 1964:
		generasi = "Baby Boomer"
	case tahunKelahiran >= 1965 && tahunKelahiran <= 1979:
		generasi = "Generasi X"
	case tahunKelahiran >= 1980 && tahunKelahiran <= 1994:
		generasi = "Generasi Y (Millennials)"
	case tahunKelahiran >= 1995 && tahunKelahiran <= 2015:
		generasi = "Generasi Z"
	default:
		generasi = "Tidak termasuk dalam kategori yang diberikan"
	}

	// Jawaban Soal Nomor 4
	fmt.Printf("Anda termasuk dalam generasi: %s\n", generasi)
}
