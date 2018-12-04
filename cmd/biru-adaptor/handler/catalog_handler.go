package handler

import (
	"context"
	"log"
	"fmt"

	"github.com/maulidihsan/interop-commerce/pkg/v1"
	"github.com/maulidihsan/interop-commerce/pkg/models"
)

func (s *server) transformCatalogRPC(ar *models.Catalog) *v1.Product {
	if ar == nil {
		return nil
	}
	res := &v1.Product{
		Id: ar.Id,
		Vendor: v1.Vendor_BLIBLI,
		Produk: ar.NamaProduk,
		Gambar:     ar.Gambar,
		Harga:   ar.Harga,
		Kategori: ar.Kategori,
		Link: ar.Url,
	}
	return res
}

func (s *server) transformCatalogData(ar *v1.Product) models.Catalog {
	res := models.Catalog{
		NamaProduk: ar.Produk,
		Url: ar.Link,
		Gambar: ar.Gambar,
		Harga: ar.Harga,
		Kategori: ar.Kategori,
	}
	return res
}

func (s *server) GetCatalog(ctx context.Context, in *v1.Keyword) (*v1.Products, error) {
	keyword := ""
	if in != nil {
		keyword = in.Keyword
	}

	list, err := s.catalog.GetCatalog(keyword)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	arrProducts := make([]*v1.Product, len(list))
	for i, a := range list {
		ar := s.transformCatalogRPC(&a)
		arrProducts[i] = ar
	}
	result := &v1.Products{
		Vendor: v1.Vendor_BLIBLI,
		Products: arrProducts,
	}
	fmt.Println("Getting Catalog: ", in.GetKeyword())
	return result, nil
}

func (s *server) GetByCategory(ctx context.Context, in *v1.Keyword) (*v1.Products, error) {
	keyword := ""
	if in != nil {
		keyword = in.Keyword
	}

	list, err := s.catalog.GetByCategory(keyword)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	arrProducts := make([]*v1.Product, len(list))
	for i, a := range list {
		ar := s.transformCatalogRPC(&a)
		arrProducts[i] = ar
	}
	result := &v1.Products{
		Vendor: v1.Vendor_BLIBLI,
		Products: arrProducts,
	}
	fmt.Println("Getting Catalog By Category: ", in.GetKeyword())
	return result, nil
}

func (s *server) UpdateCatalog(c context.Context, prod *v1.Products) (*v1.Response, error) {
	arrProducts := make([]models.Catalog, len(prod.Products))
	for i, a := range prod.Products {
		ar := s.transformCatalogData(a)
		arrProducts[i] = ar
	}
	update, err := s.catalog.UpdateCatalog(arrProducts)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	res := s.transformResponseRPC(update)
	return res, nil
}


