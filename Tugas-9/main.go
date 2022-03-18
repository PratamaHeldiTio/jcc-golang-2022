package main

import (
	lib "Tugas-9/library"
	"fmt"
)

func main() {
	// soal 1
	fmt.Println("=========Segitiga sama sisi==========")
	var segitigaSamaSisi lib.HitungBangunDatar = lib.SegitigaSamaSisi{5, 5}
	fmt.Println("Luas Segitiga sama sisi", segitigaSamaSisi.Luas())
	fmt.Println("Keliling Segitiga sama sisi", segitigaSamaSisi.Keliling())

	fmt.Println("=========Persegi panjang==========")
	var persegiPanjang lib.HitungBangunDatar = lib.PersegiPanjang{5, 5}
	fmt.Println("Luas persegi panjang", persegiPanjang.Luas())
	fmt.Println("Keliling persegi panjang", persegiPanjang.Keliling())

	fmt.Println("=========Tabung==========")
	var tabung lib.HitungBangunRuang = lib.Tabung{7, 10}
	fmt.Println("Volume tabung", tabung.Volume())
	fmt.Println("Luas permukaan tabung", tabung.LuasPermukaan())

	fmt.Println("=========Balok==========")
	var balok lib.HitungBangunRuang = lib.Balok{10, 5, 5}
	fmt.Println("Volume balok", balok.Volume())
	fmt.Println("Luas permukaan balok", balok.LuasPermukaan())

	// soal 2
	var samsung lib.ShowData = lib.Phone{
		Name:   "Samsung Galaxy Note 20",
		Brand:  "Samsung Galaxy Note 20",
		Year:   2020,
		Colors: []string{"MYstic Bronze", "MYstic White", "MYstic Black"},
	}
	fmt.Println(samsung.ShowDataPhone())

	//soal 3

	fmt.Println(lib.LuasPersegi(4, true))

	fmt.Println(lib.LuasPersegi(8, false))

	fmt.Println(lib.LuasPersegi(0, true))

	fmt.Println(lib.LuasPersegi(0, false))

	// soal 4

	var prefix interface{} = "hasil penjumlahan dari "

	var kumpulanAngkaPertama interface{} = []int{6, 8}

	var kumpulanAngkaKedua interface{} = []int{12, 14}

	// tulis jawaban anda disini
	sPrefix := prefix.(string)
	arrKumpulanAngkaPertama := kumpulanAngkaPertama.([]int)
	arrKumpulanAngkaKedua := kumpulanAngkaKedua.([]int)

	var result int

	for _, value := range arrKumpulanAngkaPertama {
		result += value
	}
	for _, value := range arrKumpulanAngkaKedua {
		result += value
	}

	fmt.Printf("%s%d + %d + %d + %d = %d",
		sPrefix,
		arrKumpulanAngkaPertama[0],
		arrKumpulanAngkaPertama[1],
		arrKumpulanAngkaKedua[0],
		arrKumpulanAngkaPertama[1],
		result)
}
