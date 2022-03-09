package main

import (
	"fmt"
	"strconv"
)

func main() {
	// soal 1
	panjangPersegiPanjang := "8"
	panjangPersegiPanjangNum, _ := strconv.Atoi(panjangPersegiPanjang)

	lebarPersegiPanjang := "5"
	lebarPersegiPanjangNum, _ := strconv.Atoi(lebarPersegiPanjang)

	alasSegitiga := "6"
	alasSegitigaNum, _ := strconv.Atoi(alasSegitiga)

	tinggiSegitiga := "7"
	tinggiSegitigaNum, _ := strconv.Atoi(tinggiSegitiga)

	kelilingPersegiPanjang := 2 * (panjangPersegiPanjangNum + lebarPersegiPanjangNum)
	luasSegitiga := alasSegitigaNum * tinggiSegitigaNum / 2 // disini tidak saya bikin 1 / 2 karena apapun dikali satu tetap nilai tersebut, dan untuk menghindari error float dikali int

	fmt.Println(kelilingPersegiPanjang)
	fmt.Println(luasSegitiga)

	// soal 2
	nilaiJohn := 80
	nilaiDoe := 50

	if nilaiJohn >= 80 {
		fmt.Println("Index nilai John A")
	} else if nilaiJohn >= 70 {
		fmt.Println("Index nilai John B")
	} else if nilaiJohn >= 60 {
		fmt.Println("Index nilai John C")
	} else if nilaiJohn >= 50 {
		fmt.Println("Index nilai John D")
	} else {
		fmt.Println("Index nilai John E")
	}

	switch {
	case nilaiDoe >= 80:
		fmt.Println("Index nilai John A")
	case nilaiDoe >= 70:
		fmt.Println("Index nilai John B")
	case nilaiDoe >= 60:
		fmt.Println("Index nilai John C")
	case nilaiDoe >= 50:
		fmt.Println("Index nilai John D")
	default:
		fmt.Println("Index nilai John E")
	}

	// soal 3
	tanggal := 9
	bulan := 9
	tahun := 2000

	switch bulan {
	case 1:
		fmt.Printf("kamu lahir pada tanggal %d Januari %d \n", tanggal, tahun)
	case 2:
		fmt.Printf("kamu lahir pada tanggal %d Februari %d \n", tanggal, tahun)
	case 3:
		fmt.Printf("kamu lahir pada tanggal %d Maret %d \n", tanggal, tahun)
	case 4:
		fmt.Printf("kamu lahir pada tanggal %d April %d \n", tanggal, tahun)
	case 5:
		fmt.Printf("kamu lahir pada tanggal %d Mei %d \n", tanggal, tahun)
	case 6:
		fmt.Printf("kamu lahir pada tanggal %d Juni %d \n", tanggal, tahun)
	case 7:
		fmt.Printf("kamu lahir pada tanggal %d Juli %d \n", tanggal, tahun)
	case 8:
		fmt.Printf("kamu lahir pada tanggal %d Agustus %d \n", tanggal, tahun)
	case 9:
		fmt.Printf("kamu lahir pada tanggal %d September %d \n", tanggal, tahun)
	case 10:
		fmt.Printf("kamu lahir pada tanggal %d Oktober %d \n", tanggal, tahun)
	case 11:
		fmt.Printf("kamu lahir pada tanggal %d November %d \n", tanggal, tahun)
	case 12:
		fmt.Printf("kamu lahir pada tanggal %d Desember %d \n", tanggal, tahun)
	default:
		fmt.Println("bulan yang anda masukan tidak sesuai")
	}

	// soal 4
	myBorn := 2000

	if myBorn >= 1944 && myBorn <= 2015 {
		if myBorn >= 1995 {
			fmt.Println("Anda Generasi Z")
		} else if myBorn >= 1980 {
			fmt.Println("Anda Generasi Y (Millenials)")
		} else if myBorn >= 1965 {
			fmt.Println("Anda Generasi X")
		} else {
			fmt.Println("Anda Generasi Baby Boomer")
		}
	} else {
		fmt.Println("Anda tidak masuk pada generasi yang disediakan")
	}

}
