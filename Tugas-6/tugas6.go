package main

import (
	"fmt"
	"math"
)

// Soal 1
var luasLingkaran float64
var kelilingLingkaran float64

func hitungLingkaran(jariJari *float64) {
	luasLingkaran = math.Pi * math.Pow(*jariJari, 2)
	kelilingLingkaran = 2 * math.Pi * *jariJari
}

// Soal 2
func introduce(sentence *string, name, gender, occupation, age string) {
	if gender == "laki-laki" {
		*sentence = fmt.Sprintf("Pak %s adalah seorang %s yang berusia %s tahun", name, occupation, age)
	} else if gender == "perempuan" {
		*sentence = fmt.Sprintf("Bu %s adalah seorang %s yang berusia %s tahun", name, occupation, age)
	}
}

// Soal 3
var buah = []string{}

func tambahBuah(buahList *[]string) {
	*buahList = append(*buahList, "Jeruk", "Semangka", "Mangga", "Strawberry", "Durian", "Manggis", "Alpukat")
}

// Soal 4
var dataFilm = []map[string]string{}

func tambahDataFilm(title, duration, genre, year string, dataFilmList *[]map[string]string) {
	film := map[string]string{
		"title":    title,
		"duration": duration,
		"genre":    genre,
		"year":     year,
	}
	*dataFilmList = append(*dataFilmList, film)
}

func main() {
	// Jawaban Nomor 1
	jariJari := 5.0
	hitungLingkaran(&jariJari)
	fmt.Printf("Luas lingkaran dengan jari-jari %.2f adalah %.2f\n", jariJari, luasLingkaran)
	fmt.Printf("Keliling lingkaran dengan jari-jari %.2f adalah %.2f\n", jariJari, kelilingLingkaran)

	// Jawaban Nomor 2
	var sentence string
	introduce(&sentence, "John", "laki-laki", "penulis", "30")
	fmt.Println(sentence)

	introduce(&sentence, "Sarah", "perempuan", "model", "28")
	fmt.Println(sentence)

	// Jawaban Nomor 3
	tambahBuah(&buah)

	for i := 0; i < len(buah); i++ {
		fmt.Printf("%d. %s\n", i+1, buah[i])
	}

	// Jawaban Nomor 4
	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	for i, film := range dataFilm {
		fmt.Printf("%d. Title: %s\n", i+1, film["title"])
		fmt.Printf("   Duration: %s\n", film["duration"])
		fmt.Printf("   Genre: %s\n", film["genre"])
		fmt.Printf("   Year: %s\n", film["year"])
		fmt.Println("---------------------------")
	}
}
