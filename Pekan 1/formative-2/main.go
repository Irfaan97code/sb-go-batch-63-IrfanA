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
	var angkaPertama = "8"
	var num, err = strconv.Atoi(angkaPertama)

	if err == nil {
		fmt.Println(num)
	}

	var angkaKedua = "5"
	var numA, errA = strconv.Atoi(angkaKedua)

	if errA == nil {
		fmt.Println(numA)
	}

	var angkaKetiga = "6"
	var numB, errB = strconv.Atoi(angkaKetiga)

	if errB == nil {
		fmt.Println(numB)
	}

	var angkaKeempat = "7"
	var numC, errC = strconv.Atoi(angkaKeempat)

	if errC == nil {
		fmt.Println(numC)
	}

	jumlah := num + numA + numB + numC
	fmt.Println(jumlah)

	//soal 5
	kalimat := `"halo halo bandung"`
	find = "halo halo"
	var kataGanti = "Hi Hi"

	var newSentence = strings.Replace(kalimat, find, kataGanti, 1)
	angka := 2021
	fmt.Println(newSentence, -angka)
}
