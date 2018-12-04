package repository

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/maulidihsan/interop-commerce/pkg/mongo"
	"github.com/maulidihsan/interop-commerce/pkg/models"
)

type UserCollection struct {
	collection *mgo.Collection
}

func NewUserCollection(session *mongo.Session, dbName string, collectionName string) *UserCollection {
	collection := session.GetCollection(dbName, collectionName)
	collection.EnsureIndex(userModelIndex())
	return &UserCollection {collection}
}

func(p *UserCollection) AddUser(u *models.User) error {
	user, err := newUserModel(u)
	if err != nil {
		return err
	}
	return p.collection.Insert(&user)
}

func(p *UserCollection) GetByEmail(email string) (*models.User,error) {
	model := userModel{}
	err := p.collection.Find(bson.M{"email": email}).One(&model)
	return model.toUser(), err
}

func(p *UserCollection) Login(c *models.Credential) (*models.User, error) {
	model := userModel{}
	err := p.collection.Find(bson.M{"email": c.Email}).One(&model)
	if err != nil {
		return nil, err
	}
	err = model.comparePassword(c.Password)
	if err != nil {
		return nil, err
	}
	return model.toUser(), nil
}
