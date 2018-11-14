package repository

import (
	// "fmt"
	"gopkg.in/mgo.v2/bson"
	"github.com/maulidihsan/flashdeal-webservice/pkg/models"
	"gopkg.in/mgo.v2"
	"github.com/maulidihsan/flashdeal-webservice/pkg/mongo"
)
type OrderCollection struct {
	collection *mgo.Collection
}

func NewOrderCollection(session *mongo.Session, dbName string, collectionName string) *OrderCollection {
	collection := session.GetCollection(dbName, collectionName)
	return &OrderCollection{collection}
}

func(p *OrderCollection) GetOrders(email string) ([]models.Order, error) {
	var models OrderArray
	err := p.collection.Find(bson.M{"pembeli.email": email}).All(&models)
	return models.toOrders(), err
}

func(p *OrderCollection) CreateOrder(n *models.Order) (*models.Response, error) {
	// fmt.Printf("%v", n)
	order := newOrderModel(n)
	err := p.collection.Insert(&order)
	if err != nil {
		return &models.Response{
			Success: false,
			Code: 400,
			Message: "failed to insert",
		}, err
	} else {
		return &models.Response{
			Success: true,
			Code: 200,
			Message: "success to insert",
		}, err
	}
}

func(p *OrderCollection) UpdateStatusOrder(id string, status string) (*models.Response, error) {
	err := p.collection.Update(bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$set": bson.M{"status": status}})
	if err != nil {
		return &models.Response{
			Success: false,
			Code: 400,
			Message: "failed to insert",
		}, err
	} else {
		return &models.Response{
			Success: true,
			Code: 200,
			Message: "success to insert",
		}, err
	}
}

func(p *OrderCollection) DeleteOrder(id string) (*models.Response, error) {
	err := p.collection.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	if err != nil {
		return &models.Response{
			Success: false,
			Code: 400,
			Message: "failed to insert",
		}, err
	} else {
		return &models.Response{
			Success: true,
			Code: 200,
			Message: "success to insert",
		}, err
	}
}