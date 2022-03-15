package main

import "fmt"

// soal 1
func hitungLingkaran(luasLingkaran *float64, kelilingLingkaran *float64, jariJari float64) {
	const PI = 3.14

	*luasLingkaran = PI * jariJari * jariJari
	*kelilingLingkaran = PI * 2 * jariJari
}

// soal 2
func introduce(sentence *string, nama string, kelamin string, profesi string, umur string) {
	if kelamin == "laki-laki" {
		*sentence = fmt.Sprintf("Pak %s adalah seorang %s yang berusia %s tahun", nama, profesi, umur)
	} else if kelamin == "perempuan" {
		*sentence = fmt.Sprintf("Bu %s adalah seorang %s yang berusia %s tahun", nama, profesi, umur)
	}
}

// soal 3
func tambahBuah(buah *[]string) {
	*buah = append(*buah, "Jeruk", "Semangka", "Mangga", "Strawberry", "Durian", "Manggis", "Alpukat")
}

// soal 4
func tambahDataFilm(title string, durasi string, genre string, release string, dataFilm *[]map[string]string) {
	temp := map[string]string{
		"genre":    genre,
		"duration": durasi,
		"year":     release,
		"title":    title,
	}

	*dataFilm = append(*dataFilm, temp)
}

func main() {

	//soal 1
	var luasLigkaran float64
	var kelilingLingkaran float64

	hitungLingkaran(&luasLigkaran, &kelilingLingkaran, 14)

	fmt.Println(luasLigkaran)
	fmt.Println(kelilingLingkaran)

	// soal 2

	var sentence string
	introduce(&sentence, "John", "laki-laki", "penulis", "30")

	fmt.Println(sentence) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"
	introduce(&sentence, "Sarah", "perempuan", "model", "28")

	fmt.Println(sentence) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	// soal 3
	var buah = []string{}

	tambahBuah(&buah)

	for index, value := range buah {
		fmt.Printf("%d. %s \n", index+1, value)
	}

	// soal 4

	var dataFilm = []map[string]string{}

	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	// isi dengan jawaban anda untuk menampilkan data
	for index, item := range dataFilm {
		fmt.Printf("%d. title : %s \n   duration : %s \n   genre : %s \n   year : %s \n\n",
			index+1, item["title"],
			item["duration"],
			item["genre"],
			item["year"])
	}
}
