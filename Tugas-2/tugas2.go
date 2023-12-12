package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Soal 1
	var kata1 = "Bootcamp"
	var kata2 = "Digital"
	var kata3 = "Skill"
	var kata4 = "Sanbercode"
	var kata5 = "Golang"

	semuakata := kata1 + " " + kata2 + " " + kata3 + " " + kata4 + " " + kata5
	// Jawaban soal 1
	fmt.Println(semuakata)

	// Soal 2
	halo := "Halo Dunia"
	halo = strings.Replace(halo, "Dunia", "Golang", 1)
	// Jawaban soal 2
	fmt.Println(halo)

	// soal3
	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"

	hasilGabungan := kataPertama + " " + kataKedua + " " + strings.Title(kataKetiga) + " " + strings.ToUpper(kataKeempat)
	// Jawaban soal 3
	fmt.Println(hasilGabungan)

	// Soal 4
	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"

	angkaPertamaInt, _ := strconv.Atoi(angkaPertama)
	angkaKeduaInt, _ := strconv.Atoi(angkaKedua)
	angkaKetigaInt, _ := strconv.Atoi(angkaKetiga)
	angkaKeempatInt, _ := strconv.Atoi(angkaKeempat)

	total := angkaPertamaInt + angkaKeduaInt + angkaKetigaInt + angkaKeempatInt

	// Jawaban Soal 4
	fmt.Println("Hasil penjumlahan:", total)

	// Soal 5
	kalimat := "halo halo bandung"
	angka := 2021

	parts := strings.Fields(kalimat)
	kalimat = "Hi Hi " + parts[2]

	// Jawaban soal 5
	fmt.Printf("%s - %d\n", kalimat, angka)

}
