package main

import (
	"fmt"
	"sync"
)

func showPhone(phones []string, wg *sync.WaitGroup) {

	for _, value := range phones {
		fmt.Println(value)
	}
	wg.Done()
}

func main() {
	// soal

	var wg sync.WaitGroup
	var phones = []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}
	wg.Add(1)
	go showPhone(phones, &wg)
	wg.Wait()

	//soal 2
	var movies = []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}

	moviesChannel := make(chan string)

	go getMovies(moviesChannel, movies...)

	for value := range moviesChannel {
		fmt.Println(value)
	}
}
