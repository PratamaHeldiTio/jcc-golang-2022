package main

import (
	"fmt"
)

// soal 1
type Buah struct {
	nama, warna string
	adaBijinya  bool
	harga       int
}

func addValueStructBuah(nama string, warna string, adaBijinya bool, harga int) Buah {
	buah := Buah{
		nama:       nama,
		warna:      warna,
		adaBijinya: adaBijinya,
		harga:      harga,
	}

	return buah
}

// soal 2
type segitiga struct {
	alas, tinggi int
}

type persegi struct {
	sisi int
}

type persegiPanjang struct {
	panjang, lebar int
}

func (s segitiga) LuasSegitiga() int {
	return s.alas * s.tinggi / 2
}
func (p persegi) LuasPersegi() int {
	return p.sisi * p.sisi
}
func (pp persegiPanjang) LuasPersegiPanjang() int {
	return pp.panjang * pp.lebar
}

// soal 3

type phone struct {
	name, brand string
	year        int
	colors      []string
}

func (p *phone) addColorPhone(color string) {
	p.colors = append(p.colors, color)
}

// soal 4
type movie struct {
	title, genre   string
	duration, year int
}

func tambahDataFilm(title string, durasi int, genre string, year int, dataFilm *[]movie) {
	temp := movie{
		genre:    genre,
		duration: durasi,
		year:     year,
		title:    title,
	}

	*dataFilm = append(*dataFilm, temp)
}

func main() {
	// soal 1
	buah1 := addValueStructBuah("Nanas", "Kuning", false, 9000)
	buah2 := addValueStructBuah("Jeruk", "Oranye", true, 8000)
	buah3 := addValueStructBuah("Semangka", "Hijau & Merah", true, 10000)
	buah4 := addValueStructBuah("Pisang", "Kuning", false, 5000)

	allBuah := []Buah{buah1, buah2, buah3, buah4}
	for _, value := range allBuah {
		if value.adaBijinya {
			fmt.Printf("nama: %s, warna: %s, ada bijinya: ada, harga: %d \n", value.nama, value.warna, value.harga)
		} else {
			fmt.Printf("nama: %s, warna: %s, ada bijinya: tidak, harga: %d \n", value.nama, value.warna, value.harga)
		}
	}

	// soal 2
	segitiga := segitiga{
		alas:   10,
		tinggi: 5,
	}
	fmt.Println(segitiga.LuasSegitiga())

	persegi := persegi{
		sisi: 5,
	}
	fmt.Println(persegi.LuasPersegi())

	persegiPanjang := persegiPanjang{
		panjang: 10,
		lebar:   5,
	}
	fmt.Println(persegiPanjang.LuasPersegiPanjang())

	// soal 3
	xiaomi := phone{
		name:  "Redmi 9",
		brand: "Xiaomi",
		year:  2019,
	}

	xiaomi.addColorPhone("red")
	xiaomi.addColorPhone("grey")
	xiaomi.addColorPhone("aqua")
	fmt.Println(xiaomi)

	// soal 4
	var dataFilm = []movie{}

	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)

	// isi dengan jawaban anda untuk menampilkan data
	for index, item := range dataFilm {
		fmt.Printf("%d. title : %v \n   duration : %v \n   genre : %v \n   year : %v \n\n",
			index+1,
			item.title,
			item.duration,
			item.genre,
			item.year,
		)
	}
}
