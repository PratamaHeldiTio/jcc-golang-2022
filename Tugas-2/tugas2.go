package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// soal 1
	firstWord := "Jabar"
	secondWord := "Coding"
	thirdWord := "Camp"
	fourthWord := "Golang"
	year := 2022
	str := "%s %s %s %s %d \n"

	fmt.Printf(str, firstWord, secondWord, thirdWord, fourthWord, year)

	// soal 2
	halo := "Halo Dunia"
	halo = strings.Replace(halo, "Dunia", "Golang", 1)

	fmt.Println(halo)

	// soal 3
	kataPertama := "saya"
	kataKedua := "senang"
	kataKetiga := "belajar"
	kataKeempat := "golang"

	kataKedua = strings.Replace(kataKedua, "s", "S", 1)
	kataKetiga = strings.Replace(kataKetiga, "r", "R", 1)
	kataKeempat = strings.ToUpper(kataKeempat)
	str3 := "%s %s %s %s \n"

	fmt.Printf(str3, kataPertama, kataKedua, kataKetiga, kataKeempat)

	// soal 4
	angkaPertama := "8"
	angkaKedua := "5"
	angkaKetiga := "6"
	angkaKeempat := "7"

	angkaPertamaNum, _ := strconv.Atoi(angkaPertama)
	angkaKeduaNum, _ := strconv.Atoi(angkaKedua)
	angkaKetigaNum, _ := strconv.Atoi(angkaKetiga)
	angkaKeempatNum, _ := strconv.Atoi(angkaKeempat)
	result := angkaPertamaNum + angkaKeduaNum + angkaKetigaNum + angkaKeempatNum
	fmt.Println(result)

	// soal 5
	kalimat := "halo halo bandung"
	kalimat = strings.Replace(kalimat, "halo", "Hi", 2)
	angka := 2022
	str5 := "\"%s\" - %d"

	fmt.Printf(str5, kalimat, angka)

}
