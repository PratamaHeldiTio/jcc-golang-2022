package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

type NilaiMahasiswa struct {
	Nama        string `json:"nama"`
	MataKuliah  string `json:"matakuliah"`
	IndeksNilai string `json:"indexNilai"`
	Nilai       uint   `json:"nilai"`
	ID          uint   `json:"id"`
}

var nilaiNilaiMahasiswa = []NilaiMahasiswa{}

// fuction for basic
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uname, pwd, ok := r.BasicAuth()
		if !ok {
			w.Write([]byte("Username atau Password tidak boleh kosong"))
			return
		}

		if uname == "admin" && pwd == "admin" {
			next.ServeHTTP(w, r)
			return
		}

		w.Write([]byte("Username atau Password tidak sesuai"))
		return
	})
}

// function for post data
func NilaiMhs(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		nilaiMahasiswa := NilaiMahasiswa{}

		if r.Header.Get("Content-Type") == "application/json" {
			// parse dari json
			decodeJson := json.NewDecoder(r.Body)
			if err := decodeJson.Decode(&nilaiMahasiswa); err != nil {
				log.Fatal(err)
			}

			if nilaiMahasiswa.Nilai > 100 {
				http.Error(w, "Nilai tidak boleh lebih dari 100", http.StatusUnauthorized)
				return
			}

			index := indexValue(nilaiMahasiswa.Nilai)
			nilaiMahasiswa.IndeksNilai = index
			nilaiMahasiswa.ID = uint(rand.Uint32())
			nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, nilaiMahasiswa)
		}

		dataNilaiNilaiMhs, _ := json.Marshal(nilaiNilaiMahasiswa)
		w.Write(dataNilaiNilaiMhs)
		return
	}

	if r.Method == "GET" {
		dataNilaiMhs, err := json.Marshal(nilaiNilaiMahasiswa)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(dataNilaiMhs)
		return
	}
	http.Error(w, "NOT FOUND", http.StatusNotFound)
	return
}

// func for indexValue
func indexValue(nilai uint) string {
	if nilai >= 80 {
		return "A"
	} else if nilai >= 70 {
		return "B"
	} else if nilai >= 60 {
		return "C"
	} else if nilai >= 50 {
		return "D"
	} else {
		return "E"
	}
}

func main() {
	http.HandleFunc("/nilai_mhs", NilaiMhs)

	fmt.Println("server running at http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
