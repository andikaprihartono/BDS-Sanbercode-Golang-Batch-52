package main

import (
	"fmt"
	"math"
)

// Soal nomor 1

type segitigaSamaSisi struct {
	alas, tinggi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type tabung struct {
	jariJari, tinggi float64
}

type balok struct {
	panjang, lebar, tinggi int
}

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

// Method untuk interface hitungBangunDatar
func (s segitigaSamaSisi) luas() int {
	return (s.alas * s.tinggi) / 2
}

func (s segitigaSamaSisi) keliling() int {
	return 3 * s.alas
}

func (p persegiPanjang) luas() int {
	return p.panjang * p.lebar
}

func (p persegiPanjang) keliling() int {
	return 2 * (p.panjang + p.lebar)
}

// Method untuk interface hitungBangunRuang
func (t tabung) volume() float64 {
	return math.Pi * math.Pow(t.jariJari, 2) * t.tinggi
}

func (t tabung) luasPermukaan() float64 {
	return 2 * math.Pi * t.jariJari * (t.jariJari + t.tinggi)
}

func (b balok) volume() float64 {
	return float64(b.panjang * b.lebar * b.tinggi)
}

func (b balok) luasPermukaan() float64 {
	return 2 * (float64(b.panjang*b.lebar) + float64(b.panjang*b.tinggi) + float64(b.lebar*b.tinggi))
}

// Soal Nomor 2
type phone struct {
	name, brand string
	year        int
	colors      []string
}

type displayData interface {
	display() string
}

func (p phone) display() string {
	return fmt.Sprintf("Phone name: %s\nBrand: %s\nYear: %d\nColors: %v\n", p.name, p.brand, p.year, p.colors)
}

// Soal Nomor 3
func luasPersegi(sisi int, display bool) interface{} {
	if sisi == 0 {
		if display {
			return "Maaf anda belum menginput sisi dari persegi"
		}
		return nil
	}
	luas := sisi * sisi
	if display {
		return fmt.Sprintf("luas persegi dengan sisi %d cm adalah %d cm", sisi, luas)
	}
	return luas
}

// Soal Nomor 4
var prefix interface{} = "hasil penjumlahan dari "
var kumpulanAngkaPertama interface{} = []int{6, 8}
var kumpulanAngkaKedua interface{} = []int{12, 14}

func main() {
	// Jawaban Nomor 1
	sSegitiga := segitigaSamaSisi{alas: 5, tinggi: 4}
	pPersegi := persegiPanjang{panjang: 6, lebar: 3}
	tTabung := tabung{jariJari: 2.5, tinggi: 5}
	bBalok := balok{panjang: 4, lebar: 3, tinggi: 2}

	fmt.Printf("Luas segitiga sama sisi: %d\n", sSegitiga.luas())
	fmt.Printf("Keliling segitiga sama sisi: %d\n", sSegitiga.keliling())
	fmt.Printf("Luas persegi panjang: %d\n", pPersegi.luas())
	fmt.Printf("Keliling persegi panjang: %d\n", pPersegi.keliling())
	fmt.Printf("Volume tabung: %.2f\n", tTabung.volume())
	fmt.Printf("Luas permukaan tabung: %.2f\n", tTabung.luasPermukaan())
	fmt.Printf("Volume balok: %.2f\n", bBalok.volume())
	fmt.Printf("Luas permukaan balok: %.2f\n", bBalok.luasPermukaan())

	// Jawaban Nomor 2
	p := phone{
		name:   "iPhone 12",
		brand:  "Apple",
		year:   2020,
		colors: []string{"Black", "White"},
	}

	fmt.Println(p.display())

	// Jawaban Nomor 3
	fmt.Println(luasPersegi(4, true))
	fmt.Println(luasPersegi(8, false))
	fmt.Println(luasPersegi(0, true))
	fmt.Println(luasPersegi(0, false))

	// Jawaban Nomor 4
	sum := 0
	switch v := kumpulanAngkaPertama.(type) {
	case []int:
		for _, num := range v {
			sum += num
		}
	}

	switch v := kumpulanAngkaKedua.(type) {
	case []int:
		for _, num := range v {
			sum += num
		}
	}

	fmt.Printf("%s%d + %d + %d + %d = %d\n", prefix, kumpulanAngkaPertama.([]int)[0], kumpulanAngkaPertama.([]int)[1], kumpulanAngkaKedua.([]int)[0], kumpulanAngkaKedua.([]int)[1], sum)

}
