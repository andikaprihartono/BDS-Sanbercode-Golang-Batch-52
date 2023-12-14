package main

import "fmt"

func main() {
	// soal nomor 1
	totalIterasi := 20

	for i := 1; i <= totalIterasi; i++ {
		// Check apakah angka ganjil atau genap
		if i%2 == 1 {
			fmt.Printf("%d - Santai\n", i)
		} else {
			fmt.Printf("%d - Berkualitas\n", i)
		}

		// jawaban nomor 1
		if i%3 == 0 && i%2 == 1 {
			fmt.Printf("%d - I Love Coding\n", i)
		}
	}

	// soal nomor 2
	tinggi := 7

	for i := 1; i <= tinggi; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("#")
		}
		// Jawaban Nomor 2
		fmt.Println()
	}

	// Soal Nomor 3
	kalimat := [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}

	hasilOlah := fmt.Sprintf("[%s]", kalimat[2]+" "+kalimat[3]+" "+kalimat[4]+" "+kalimat[5]+" "+kalimat[6])

	// Jawaban Nomor 3
	fmt.Println(hasilOlah)

	// Soal Nomor 4
	sayuran := []string{}
	sayuran = append(sayuran, "Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun")

	// Jawaban Nomor 4
	for i, s := range sayuran {
		fmt.Printf("%d. %s\n", i+1, s)
	}

	// Soal Nomor 5
	satuan := map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}

	// Jawaban Soal Nomor 5
	for k, v := range satuan {
		fmt.Printf("%s: %d\n", k, v)
	}

	volume := satuan["panjang"] * satuan["lebar"] * satuan["tinggi"]
	fmt.Printf("Volume Balok: %d\n", volume)
}
