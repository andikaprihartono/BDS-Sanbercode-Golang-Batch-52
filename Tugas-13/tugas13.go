package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type NilaiMahasiswa struct {
	Nama        string `json:"nama"`
	MataKuliah  string `json:"mata_kuliah"`
	IndeksNilai string `json:"indeks_nilai"`
	Nilai       uint   `json:"nilai"`
	ID          uint   `json:"id"`
}

var nilaiNilaiMahasiswa = []NilaiMahasiswa{}
var users = map[string]string{
	"username": "password", // contoh: username dan password
}

func calculateIndeksNilai(nilai uint) string {
	switch {
	case nilai >= 80:
		return "A"
	case nilai >= 70:
		return "B"
	case nilai >= 60:
		return "C"
	case nilai >= 50:
		return "D"
	default:
		return "E"
	}
}

func addNilaiMahasiswaHandler(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	if !ok || users[username] != password {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var nilaiData NilaiMahasiswa

	if r.Header.Get("Content-Type") == "application/json" {
		err := json.NewDecoder(r.Body).Decode(&nilaiData)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
	} else {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Invalid form data", http.StatusBadRequest)
			return
		}
		nilaiData.Nama = r.FormValue("nama")
		nilaiData.MataKuliah = r.FormValue("mata_kuliah")
		value, err := strconv.Atoi(r.FormValue("nilai"))
		if err != nil || value > 100 {
			http.Error(w, "Invalid value for Nilai", http.StatusBadRequest)
			return
		}
		nilaiData.Nilai = uint(value)
	}

	nilaiData.IndeksNilai = calculateIndeksNilai(nilaiData.Nilai)
	// Anda dapat mengatur ID dengan cara tertentu sesuai kebutuhan aplikasi Anda.
	nilaiData.ID = uint(len(nilaiNilaiMahasiswa) + 1)

	nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, nilaiData)
	json.NewEncoder(w).Encode(nilaiData)
}

func getNilaiMahasiswaHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(nilaiNilaiMahasiswa)
}

func main() {
	http.HandleFunc("/addNilai", addNilaiMahasiswaHandler)
	http.HandleFunc("/getNilai", getNilaiMahasiswaHandler)

	fmt.Println("Server started at localhost:8080")
	http.ListenAndServe(":8080", nil)
}
