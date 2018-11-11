package models

type User struct {
	Id string `json:"id"`
	Email string `json:"email"`
	Password string `json:"password"`
	Nama string `json:"nama"`
	Alamat string `json:"alamat"`
	Telepon string `json:"telepon"`
}

type UserService interface {
	AddUser(u *User) error
	GetByEmail(email string) (*User,error)
}