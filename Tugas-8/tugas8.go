package main

import (
	"fmt"
	"math"
	"strings"
)

// soal 1

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
	panjang, lebar, tinggi float64
}

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

func (sss segitigaSamaSisi) luas() int {
	return sss.alas * sss.tinggi / 2
}

func (sss segitigaSamaSisi) keliling() int {
	return sss.alas * 3
}

func (pp persegiPanjang) luas() int {
	return pp.panjang * pp.lebar
}

func (pp persegiPanjang) keliling() int {
	return 2 * (pp.panjang + pp.lebar)
}

func (t tabung) volume() float64 {
	return math.Pi * math.Pow(t.jariJari, 2) * t.tinggi
}

func (t tabung) luasPermukaan() float64 {
	return (2 * math.Pi * t.jariJari * t.tinggi) + (math.Pi * math.Pow(t.jariJari, 2))
}

func (b balok) volume() float64 {
	return b.panjang * b.lebar * b.tinggi
}

func (b balok) luasPermukaan() float64 {
	return 2 * (b.panjang*b.lebar + b.panjang*b.tinggi + b.lebar*b.tinggi)
}

// soal 2

type phone struct {
	name, brand string
	year        int
	colors      []string
}

type showData interface {
	showDataPhone() string
}

func (phone phone) showDataPhone() string {
	data := fmt.Sprintf("name : %s \nbrand : %s \nyear : %d \ncolors : %s", phone.name, phone.brand, phone.year, strings.Join(phone.colors, ", "))

	return data
}

// soal 3
type persegi interface{}

func luasPersegi(sisi int, kondisi bool) persegi {
	luas := sisi * sisi
	var result persegi
	if sisi > 0 {
		if kondisi {
			result = fmt.Sprintf("luas persegi dengan sisi %d cm adalah %d cm", sisi, luas)
		} else {
			result = luas
		}
	} else {
		if kondisi {
			result = fmt.Sprintf("Maaf anda belum menginput sisi dari persegi")
		} else {
			result = nil
		}
	}

	return result
}

func main() {
	// soal 1
	fmt.Println("=========Segitiga sama sisi==========")
	var segitigaSamaSisi hitungBangunDatar = segitigaSamaSisi{5, 5}
	fmt.Println("Luas Segitiga sama sisi", segitigaSamaSisi.luas())
	fmt.Println("Keliling Segitiga sama sisi", segitigaSamaSisi.keliling())

	fmt.Println("=========Persegi panjang==========")
	var persegiPanjang hitungBangunDatar = persegiPanjang{5, 5}
	fmt.Println("Luas persegi panjang", persegiPanjang.luas())
	fmt.Println("Keliling persegi panjang", persegiPanjang.keliling())

	fmt.Println("=========Tabung==========")
	var tabung hitungBangunRuang = tabung{7, 10}
	fmt.Println("Volume tabung", tabung.volume())
	fmt.Println("Luas permukaan tabung", tabung.luasPermukaan())

	fmt.Println("=========Balok==========")
	var balok hitungBangunRuang = balok{10, 5, 5}
	fmt.Println("Volume balok", balok.volume())
	fmt.Println("Luas permukaan balok", balok.luasPermukaan())

	// soal 2
	var samsung showData = phone{
		name:   "Samsung Galaxy Note 20",
		brand:  "Samsung Galaxy Note 20",
		year:   2020,
		colors: []string{"MYstic Bronze", "MYstic White", "MYstic Black"},
	}
	fmt.Println(samsung.showDataPhone())

	//soal 3

	fmt.Println(luasPersegi(4, true))

	fmt.Println(luasPersegi(8, false))

	fmt.Println(luasPersegi(0, true))

	fmt.Println(luasPersegi(0, false))

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
