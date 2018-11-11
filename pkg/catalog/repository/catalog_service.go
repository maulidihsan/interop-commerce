package repository

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/maulidihsan/flashdeal-webservice/pkg/mongo"
	"github.com/maulidihsan/flashdeal-webservice/pkg/models"
)

type CatalogCollection struct {
	collection *mgo.Collection
}

func NewCatalogCollection(session *mongo.Session, dbName string, collectionName string) *CatalogCollection {
	collection := session.GetCollection(dbName, collectionName)
	collection.EnsureIndex(catalogModelIndex())
	return &CatalogCollection{collection}
}

func(p *CatalogCollection) GetCatalog(keyword string) ([]models.Catalog,error) {
	var models CatalogArray
	err := p.collection.Find(bson.M{"product_name": bson.M{"$regex": bson.RegEx{fmt.Sprintf(".*%s.*", keyword), ""}}}).All(&models)
	//fmt.Printf("%v", models)
	return models.toCatalogs(), err
}

func(p *CatalogCollection) UpdateCatalog(c []models.Catalog) (*models.Response, error) {
	var productToInsert CatalogArray
	for _, product := range c {
		productToInsert = append(productToInsert, newCatalogModel(&product))
	}
	err := p.collection.Insert(productToInsert)
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
