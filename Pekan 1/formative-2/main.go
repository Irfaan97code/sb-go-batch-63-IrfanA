package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//soal 1
	var wordFirst string = "Bootcamp"
	wordSecond := "Digital"
	wordThird := "Skill"
	wordFourth := "Sanbercode"
	wordFiveth := "Golang"
	fmt.Println(wordFirst, wordSecond, wordThird, wordFourth, wordFiveth)

	//soal 2
	halo := "halo Dunia"
	find := "Dunia"
	var replaceWith = "Golang"

	var newText = strings.Replace(halo, find, replaceWith, 1)
	fmt.Println(newText)

	//soal 3
	var kataPertama string = "saya"
	var str = strings.ToLower(kataPertama)

	var kataKedua = "senang"
	var strA = strings.Title(kataKedua)

	var kataKetiga = "belajar"
	var strB = strings.Title(kataKetiga)

	var kataKeempat = "golang"
	var strC = strings.ToUpper(kataKeempat)

	fmt.Println(str, strA, strB, strC)

	//soal 4
	// variabel
	var angkaPertama string = "8"
	var angkaKedua string = "5"
	var angkaKetiga string = "6"
	var angkaKeempat string = "7"

	numA, _ := strconv.Atoi(angkaPertama)
	numB, _ := strconv.Atoi(angkaKedua)
	numC, _ := strconv.Atoi(angkaKetiga)
	numD, _ := strconv.Atoi(angkaKeempat)

	var jumlah = (numA + numB + numC + numD)
	fmt.Println("jumlah keseluruhan", jumlah)

	//soal 5
	kalimat := `"halo halo bandung"`
	find = "halo halo"
	var kataGanti = "Hi Hi"

	var newSentence = strings.Replace(kalimat, find, kataGanti, 1)
	angka := 2021
	fmt.Println(newSentence, -angka)
}
