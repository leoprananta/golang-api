package models

type Mahasiswa struct{
	NIM string `json:"nim" gorm:"primary_key"`
	Nama string `json:"nama"`
}