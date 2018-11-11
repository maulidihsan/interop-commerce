package user

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/maulidihsan/flashdeal-webservice/pkg/mongo"
	"github.com/maulidihsan/flashdeal-webservice/pkg/models"
)

type UserService struct {
	collection *mgo.Collection
	hash models.Hash
}

func NewUserService(session *mongo.Session, dbName string, collectionName string, hash models.Hash) *UserService {
	collection := session.GetCollection(dbName, collectionName)
	collection.EnsureIndex(userModelIndex())
	return &UserService {collection, hash}
}

func(p *UserService) AddUser(u *models.User) error {
	user := newUserModel(u)
	hashedPassword, err := p.hash.Generate(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return p.collection.Insert(&user)
}

func(p *UserService) GetByEmail(email string) (*models.User,error) {
	model := userModel{}
	err := p.collection.Find(bson.M{"email": email}).One(&model)
	return model.toUser(), err
}