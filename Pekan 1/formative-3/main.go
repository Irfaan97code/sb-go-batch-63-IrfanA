package main

import (
	"fmt"
	"strconv"
)

func main() {
	//soal 1
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"
	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	panjang, _ := strconv.Atoi(panjangPersegiPanjang)
	lebar, _ := strconv.Atoi(lebarPersegiPanjang)
	alas, _ := strconv.Atoi(alasSegitiga)
	tinggi, _ := strconv.Atoi(tinggiSegitiga)

	var luasPersegiPanjang int = panjang * lebar
	var kelilingPersegiPanjang int = 2 * (panjang + lebar)
	var luasSegitiga int = (alas * tinggi) / 2

	fmt.Println("Luas Persegi Panjang:", luasPersegiPanjang)
	fmt.Println("Keliling Persegi Panjang:", kelilingPersegiPanjang)
	fmt.Println("Luas Segitiga:", luasSegitiga)

	// soal 2
	// Variabel nilai
	var nilaiJohn = 80
	var nilaiDoe = 50

	determineIndex := func(nilai int) string {
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
	indeksJohn := determineIndex(nilaiJohn)
	indeksDoe := determineIndex(nilaiDoe)

	fmt.Println("Indeks nilai John:", indeksJohn)
	fmt.Println("Indeks nilai Doe:", indeksDoe)

	// soal 3
	// Variabel tanggal lahir
	var tanggal = 9
	var bulan = 11
	var tahun = 1997

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
	fmt.Printf("%d %s %d\n", tanggal, namaBulan, tahun)

	// soal 4
	var generasi string
	switch {
	case tahun >= 1944 && tahun <= 1964:
		generasi = "Baby Boomer"
	case tahun >= 1965 && tahun <= 1979:
		generasi = "Generasi X"
	case tahun >= 1980 && tahun <= 1994:
		generasi = "Generasi Y (Millenials)"
	case tahun >= 1995 && tahun <= 2015:
		generasi = "Generasi Z"
	default:
		generasi = "Generasi tidak terdefinisi"
	}
	fmt.Println("Anda termasuk:", generasi)
}
