package user

import (
	"github.com/maulidihsan/flashdeal-webservice/pkg/models"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

type userModel struct {
	Id bson.ObjectId `bson:"_id,omitempty"`
	Email string `bson:"email"`
	Password string `bson:"password"`
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

func newUserModel(u *models.User) *userModel {
	return &userModel{
		Email: u.Email,
		Password: u.Password,
		Nama: u.Nama,
		Alamat: u.Alamat,
		Telepon: u.Telepon,
	}
}

func(u *userModel) toUser() *models.User {
	return &models.User{
		Id: u.Id.Hex(),
		Email: u.Email,
		Password: u.Password,
		Nama: u.Nama,
		Alamat: u.Alamat,
		Telepon: u.Telepon,
	}
}