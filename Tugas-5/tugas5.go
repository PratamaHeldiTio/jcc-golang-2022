package main

import (
	"fmt"
)

// function soal 1
func luasPersegiPanjang(panjang int, lebar int) (result int) {
	result = panjang * lebar
	return
}

func kelilingPersegiPanjang(panjang int, lebar int) (result int) {
	result = 2 * (panjang + lebar)
	return
}

func volumeBalok(panjang int, lebar int, tinggi int) (result int) {
	result = panjang * lebar * tinggi
	return
}

// function soal 2
func introduce(nama string, kelamin string, profesi string, umur string) (result string) {
	if kelamin == "laki-laki" {
		result = fmt.Sprintf("Pak %s adalah seorang %s yang berusia %s tahun", nama, profesi, umur)
	} else if kelamin == "perempuan" {
		result = fmt.Sprintf("Bu %s adalah seorang %s yang berusia %s tahun", nama, profesi, umur)
	}

	return
}

// function soal 3
func buahFavorit(nama string, buah ...string) (result string) {
	panjangBuah := len(buah)
	result = fmt.Sprintf("halo nama saya %s dan buah favorit saya adalah", nama)
	for index, buah := range buah {
		if index == (panjangBuah - 1) {
			result += fmt.Sprintf(" \"%s\"", buah)
		} else {
			result += fmt.Sprintf(" \"%s\",", buah)
		}
	}

	return
}

func main() {
	// soal 1
	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println(luas)
	fmt.Println(keliling)
	fmt.Println(volume)

	// soal 2
	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"

	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	// soal 3
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}

	var buahFavoritJohn = buahFavorit("john", buah...)

	fmt.Println(buahFavoritJohn)
	// halo nama saya john dan buah favorit saya adalah "semangka", "jeruk", "melon", "pepaya"

	// soal 4
	var dataFilm = []map[string]string{}
	tambahDataFilm := func(title string, durasi string, genre string, release string) {
		temp := map[string]string{
			"genre": genre,
			"jam":   durasi,
			"tahun": release,
			"title": title,
		}

		dataFilm = append(dataFilm, temp)
	}

	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}
}
