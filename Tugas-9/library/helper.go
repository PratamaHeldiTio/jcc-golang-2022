package library

import (
	"fmt"
	"math"
	"strings"
)

// soal 1

type SegitigaSamaSisi struct {
	Alas, Tinggi int
}

type PersegiPanjang struct {
	Panjang, Lebar int
}

type Tabung struct {
	JariJari, Tinggi float64
}

type Balok struct {
	Panjang, Lebar, Tinggi float64
}

type HitungBangunDatar interface {
	Luas() int
	Keliling() int
}

type HitungBangunRuang interface {
	Volume() float64
	LuasPermukaan() float64
}

type Phone struct {
	Name, Brand string
	Year        int
	Colors      []string
}

type ShowData interface {
	ShowDataPhone() string
}

// soal 3
type Persegi interface{}

func (sss SegitigaSamaSisi) Luas() int {
	return sss.Alas * sss.Tinggi / 2
}

func (sss SegitigaSamaSisi) Keliling() int {
	return sss.Alas * 3
}

func (pp PersegiPanjang) Luas() int {
	return pp.Panjang * pp.Lebar
}

func (pp PersegiPanjang) Keliling() int {
	return 2 * (pp.Panjang + pp.Lebar)
}

func (t Tabung) Volume() float64 {
	return math.Pi * math.Pow(t.JariJari, 2) * t.Tinggi
}

func (t Tabung) LuasPermukaan() float64 {
	return (2 * math.Pi * t.JariJari * t.Tinggi) + (math.Pi * math.Pow(t.JariJari, 2))
}

func (b Balok) Volume() float64 {
	return b.Panjang * b.Lebar * b.Tinggi
}

func (b Balok) LuasPermukaan() float64 {
	return 2 * (b.Panjang*b.Lebar + b.Panjang*b.Tinggi + b.Lebar*b.Tinggi)
}

func (phone Phone) ShowDataPhone() string {
	data := fmt.Sprintf("name : %s \nbrand : %s \nyear : %d \ncolors : %s", phone.Name, phone.Brand, phone.Year, strings.Join(phone.Colors, ", "))

	return data
}

func LuasPersegi(sisi int, kondisi bool) Persegi {
	luas := sisi * sisi
	var result Persegi
	if sisi > 0 {
		if kondisi {
			result = fmt.Sprintf("luas persegi dengan sisi %d cm adalah %d cm", sisi, luas)
		} else {
			result = luas
		}
	} else {
		if kondisi {
			result = "Maaf anda belum menginput sisi dari persegi"
		} else {
			result = nil
		}
	}

	return result
}
