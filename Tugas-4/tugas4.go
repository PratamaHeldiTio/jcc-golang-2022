package main

import (
	"fmt"
)

func main() {
	// soal 1
	for i := 1; i <= 20; i++ {
		if i%2 == 1 {
			if i%3 == 0 {
				fmt.Printf("%d - I Love Coding \n", i)
			} else {
				fmt.Printf("%d - JCC \n", i)
			}
		} else {
			fmt.Printf("%d - Candradimuka \n", i)
		}
	}

	// soal 2
	for i := 0; i < 7; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("#")
		}
		fmt.Println()
	}

	// soal 3
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	fmt.Println(kalimat[2:])

	// soal 4
	var sayuran = []string{}
	sayuran = append(sayuran, "Bayam", "Buncis", "Kangkung", "Kubisa", "Seledri", "Tauge", "Timun")

	for i, sayur := range sayuran {
		fmt.Printf("%d. %s \n", i+1, sayur)
	}

	// soal 5
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}

	volumeBalok := 1
	for i, sat := range satuan {
		fmt.Printf("%s = %d \n", i, sat)
		volumeBalok *= sat
	}

	fmt.Println("Volume Balok =", volumeBalok)

}
