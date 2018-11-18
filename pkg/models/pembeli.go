package models

type Person struct {
	Id string `json:"id" validate:"omitempty"`
	Email string `json:"email" validate:"email,required"`
	Nama string `json:"nama" validate:"required"`
	Alamat string `json:"alamat" validate:"required"`
	Telepon string `json:"telepon" validate:"required"`
}