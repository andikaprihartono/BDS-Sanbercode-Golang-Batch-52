package main

import (
	"fmt"
	"tugas9/tugas9"
)

func main() {
	triangle := tugas9.SegitigaSamaSisi{Alas: 5, Tinggi: 4}
	rectangle := tugas9.PersegiPanjang{Panjang: 6, Lebar: 3}
	cylinder := tugas9.Tabung{JariJari: 2.5, Tinggi: 5}
	rectangularBox := tugas9.Balok{Panjang: 4, Lebar: 3, Tinggi: 2}

	fmt.Printf("Luas segitiga sama sisi: %d\n", triangle.Luas())
	fmt.Printf("Keliling segitiga sama sisi: %d\n", triangle.Keliling())
	fmt.Printf("Luas persegi panjang: %d\n", rectangle.Luas())
	fmt.Printf("Keliling persegi panjang: %d\n", rectangle.Keliling())
	fmt.Printf("Volume tabung: %.2f\n", cylinder.Volume())
	fmt.Printf("Luas permukaan tabung: %.2f\n", cylinder.LuasPermukaan())
	fmt.Printf("Volume balok: %.2f\n", rectangularBox.Volume())
	fmt.Printf("Luas permukaan balok: %.2f\n", rectangularBox.LuasPermukaan())
}
