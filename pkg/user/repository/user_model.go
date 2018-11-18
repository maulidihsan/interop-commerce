package repository

import (
	"github.com/maulidihsan/flashdeal-webservice/pkg/models"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type userModel struct {
	Id bson.ObjectId `bson:"_id,omitempty"`
	Email string `bson:"email"`
	PasswordHash string `bson:"password_hash"`
	Salt string `bson:"password_salt"`
	Nama string `bson:"nama"`
	Alamat string `bson:"alamat"`
	Telepon string `bson:"telepon"`
}

func userModelIndex() mgo.Index {
	return mgo.Index{
		Key: []string{"email"},
		Unique: true,
		DropDups: true,
		Background: true,
		Sparse: true,
	}
}

func newUserModel(u *models.User) (*userModel, error) {
	user := userModel{
		Email: u.Email,
		Nama: u.Nama,
		Alamat: u.Alamat,
		Telepon: u.Telepon,
	}
	err := user.setSaltedPassword(u.Password)
	return &user, err
}

func(u *userModel) comparePassword(password string) error { 
	incoming := []byte(password+u.Salt)
	existing := []byte(u.PasswordHash)
	err := bcrypt.CompareHashAndPassword(existing, incoming)
	return err
}

func(u *userModel) setSaltedPassword(password string) error { 
	salt := uuid.New().String()
	passwordBytes := []byte(password + salt)
	hash, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
    	return err
	}

	u.PasswordHash = string(hash[:])
	u.Salt = salt
	return nil
}

func(u *userModel) toUser() *models.User {
	return &models.User{
		Id: u.Id.Hex(),
		Email: u.Email,
		Nama: u.Nama,
		Alamat: u.Alamat,
		Telepon: u.Telepon,
	}
}