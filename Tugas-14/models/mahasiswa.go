package models

type NilaiMahasiswa struct {
	Nama        string `json:"nama"`
	MataKuliah  string `json:"matakuliah"`
	IndeksNilai string `json:"indexNilai"`
	Nilai       uint   `json:"nilai"`
	ID          uint   `json:"id"`
}
