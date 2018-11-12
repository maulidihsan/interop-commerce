package models

type Person struct {
	Id string `json:"id"`
	Email string `json:"email"`
	Nama string `json:"nama"`
	Alamat string `json:"alamat"`
	Telepon string `json:"telepon"`
}