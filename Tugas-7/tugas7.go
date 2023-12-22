package main

import "fmt"

// Soal Nomor 1
type buah struct {
	nama       string
	warna      string
	adaBijinya bool
	harga      int
}

// Soal Nomor 2
type segitiga struct {
	alas, tinggi int
}

type persegi struct {
	sisi int
}

type persegiPanjang struct {
	panjang, lebar int
}

// Method luas untuk struct segitiga
func (s segitiga) luas() float64 {
	return 0.5 * float64(s.alas) * float64(s.tinggi)
}

// Method luas untuk struct persegi
func (p persegi) luas() int {
	return p.sisi * p.sisi
}

// Method luas untuk struct persegiPanjang
func (pp persegiPanjang) luas() int {
	return pp.panjang * pp.lebar
}

// Soal 3
type phone struct {
	name, brand string
	year        int
	colors      []string
}

func (p *phone) addColor(color string) {
	p.colors = append(p.colors, color)
}

// Soal 4
type movie struct {
	title    string
	duration int
	genre    string
	year     int
}

// Deklarasi slice
var dataFilm = []movie{}

func tambahDataFilm(title string, duration int, genre string, year int, dataFilmList *[]movie) {
	film := movie{
		title:    title,
		duration: duration,
		genre:    genre,
		year:     year,
	}
	*dataFilmList = append(*dataFilmList, film)
}

func main() {
	// Jawaban Soal 1
	nanas := buah{
		nama:       "Nanas",
		warna:      "Kuning",
		adaBijinya: false,
		harga:      9000,
	}

	jeruk := buah{
		nama:       "Jeruk",
		warna:      "Oranye",
		adaBijinya: true,
		harga:      8000,
	}

	semangka := buah{
		nama:       "Semangka",
		warna:      "Hijau & Merah",
		adaBijinya: true,
		harga:      10000,
	}

	pisang := buah{
		nama:       "Pisang",
		warna:      "Kuning",
		adaBijinya: false,
		harga:      5000,
	}

	// Print semua objek
	fmt.Println("Nama:", nanas.nama)
	fmt.Println("Warna:", nanas.warna)
	fmt.Println("Ada bijinya:", nanas.adaBijinya)
	fmt.Println("Harga:", nanas.harga)
	fmt.Println()

	fmt.Println("Nama:", jeruk.nama)
	fmt.Println("Warna:", jeruk.warna)
	fmt.Println("Ada bijinya:", jeruk.adaBijinya)
	fmt.Println("Harga:", jeruk.harga)
	fmt.Println()

	fmt.Println("Nama:", semangka.nama)
	fmt.Println("Warna:", semangka.warna)
	fmt.Println("Ada bijinya:", semangka.adaBijinya)
	fmt.Println("Harga:", semangka.harga)
	fmt.Println()

	fmt.Println("Nama:", pisang.nama)
	fmt.Println("Warna:", pisang.warna)
	fmt.Println("Ada bijinya:", pisang.adaBijinya)
	fmt.Println("Harga:", pisang.harga)

	// Jawaban Nomor 2
	// Membuat objek dari struct
	segitiga1 := segitiga{alas: 4, tinggi: 5}
	persegi1 := persegi{sisi: 6}
	persegiPanjang1 := persegiPanjang{panjang: 7, lebar: 8}

	fmt.Printf("Luas segitiga dengan alas %d dan tinggi %d adalah: %.2f\n", segitiga1.alas, segitiga1.tinggi, segitiga1.luas())
	fmt.Printf("Luas persegi dengan sisi %d adalah: %d\n", persegi1.sisi, persegi1.luas())
	fmt.Printf("Luas persegi panjang dengan panjang %d dan lebar %d adalah: %d\n", persegiPanjang1.panjang, persegiPanjang1.lebar, persegiPanjang1.luas())

	// JAwaban Soal 3
	p := phone{
		name:  "iPhone",
		brand: "Apple",
		year:  2020,
		colors: []string{
			"Black",
			"White",
		},
	}

	// Menambahkan warna baru ke objek phone menggunakan method addColor
	p.addColor("Gold")
	p.addColor("Silver")

	// Menampilkan warna-warna dari objek phone
	fmt.Println("Colors:", p.colors)

	// JAwaban Soal 4
	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)

	// Menampilkan data dari slice dataFilm
	for i, film := range dataFilm {
		fmt.Printf("%d. Title: %s\n", i+1, film.title)
		fmt.Printf("   Duration: %d minutes\n", film.duration)
		fmt.Printf("   Genre: %s\n", film.genre)
		fmt.Printf("   Year: %d\n", film.year)
		fmt.Println("---------------------------")
	}
}
