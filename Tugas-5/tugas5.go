package main

import (
	"fmt"
	"strings"
)

// Soal Nomor 4
var dataFilm = []map[string]string{}

func tambahDataFilm() func(string, string, string, string) {
	return func(genre, jam, tahun, title string) {
		film := map[string]string{
			"Title": title,
			"Jam":   jam,
			"Genre": genre,
			"Tahun": tahun,
		}
		dataFilm = append(dataFilm, film)
	}
}

// Soal Nomor 1
func luasPersegiPanjang(panjang, lebar float64) float64 {
	return panjang * lebar
}

func kelilingPersegiPanjang(panjang, lebar float64) float64 {
	return 2 * (panjang + lebar)
}

func volumeBalok(panjang, lebar, tinggi float64) float64 {
	return panjang * lebar * tinggi
}

// Soal Nomor 2
func introduce(nama, jenisKelamin, pekerjaan, usia string) string {
	var title string

	if jenisKelamin == "laki-laki" {
		title = "Pak"
	} else if jenisKelamin == "perempuan" {
		title = "Bu"
	}

	hasil := fmt.Sprintf("%s %s adalah seorang %s yang berusia %s tahun", title, nama, pekerjaan, usia)
	return hasil
}

// Soal Nomor 3
func buahFavorit(nama string, buah ...string) string {
	buahFavoritStr := strings.Join(buah, "\", \"")
	return fmt.Sprintf("Halo, nama saya %s dan buah favorit saya adalah \"%s\"", nama, buahFavoritStr)
}

func main() {
	// Jawaban Nomor 1
	panjang := 12.0
	lebar := 4.0
	tinggi := 8.0

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println("Luas Persegi Panjang:", luas)
	fmt.Println("Keliling Persegi Panjang:", keliling)
	fmt.Println("Volume Balok:", volume)

	// Jawaban Nomor 2
	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john)

	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah)

	// Jawaban Nomor 3
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}
	buahFavoritJohn := buahFavorit("John", buah...)

	fmt.Println(buahFavoritJohn)

	// Jawaban Nomor 4
	tambahData := tambahDataFilm()

	tambahData("action", "2 jam", "1999", "LOTR")
	tambahData("action", "2 jam", "2019", "Avenger")
	tambahData("action", "2 jam", "2004", "Spiderman")
	tambahData("horror", "2 jam", "2004", "Juon")

	for _, item := range dataFilm {
		fmt.Println(item)
	}

}
