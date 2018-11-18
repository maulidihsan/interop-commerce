package models

type User struct {
	Id string `json:"id,omitempty"`
	Email string `json:"email" validate:"required,email"` 
	Password string `json:"password,omitempty" validate:"required"`
	Nama string `json:"nama" validate:"required"`
	Alamat string `json:"alamat" validate:"required"`
	Telepon string `json:"telepon" validate:"required,numeric"`
}

type Credential struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
