package handler

import (
	"log"
	"context"
	"fmt"

	"google.golang.org/grpc"
	catalogService "github.com/maulidihsan/flashdeal-webservice/pkg/catalog/usecase"
	orderService "github.com/maulidihsan/flashdeal-webservice/pkg/order/usecase"
	"github.com/maulidihsan/flashdeal-webservice/pkg/v1"
	"github.com/maulidihsan/flashdeal-webservice/pkg/models"
)

func NewCatalogServerGrpc(gserver *grpc.Server, catalogUsecase catalogService.CatalogUsecase) {

	catalogServer := &server{
		catalog: catalogUsecase,
	}
	v1.RegisterCatalogServiceServer(gserver, catalogServer)
}

func NewOrderServerGrpc(gserver *grpc.Server, orderUsecase orderService.OrderUsecase) {

	orderServer := &server{
		order: orderUsecase,
	}

	v1.RegisterOrderServiceServer(gserver, orderServer)
}

type server struct {
	catalog catalogService.CatalogUsecase
	order orderService.OrderUsecase
}

func (s *server) transformCatalogRPC(ar *models.Catalog) *v1.Product {
	if ar == nil {
		return nil
	}
	res := &v1.Product{
		Vendor: 0,
		Produk: ar.NamaProduk,
		Gambar:     ar.Gambar,
		Harga:   ar.Harga,
		Kategori: ar.Kategori,
		Link: ar.Url,
	}
	return res
}

func (s *server) transformResponseRPC(ar *models.Response) *v1.Response {
	if ar == nil {
		return nil
	}
	res := &v1.Response{
		Success: ar.Success,
		Code: ar.Code,
		Message: ar.Message,
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

func (s *server) transformOrderRPC(in *models.Order) *v1.Order {
	if in == nil {
		return nil
	}
	res := &v1.Order{
		Pembeli: &v1.Person{
			Nama: in.Pembeli.Nama,
			Alamat: in.Pembeli.Alamat,
			Telepon: in.Pembeli.Telepon,
			Email: in.Pembeli.Email,
		},
		Barang: &v1.Product{
			Produk: in.Produk.NamaProduk,
			Link: in.Produk.Url,
			Gambar: in.Produk.Gambar,
			Harga: in.Produk.Harga,
			Kategori: in.Produk.Kategori,
		},
		Status: in.Status,
	}
	return res
}

func (s *server) transformOrderData(in *v1.Order) models.Order{

	res := models.Order{
		Pembeli: models.User{
			Nama: in.Pembeli.Nama,
			Alamat: in.Pembeli.Alamat,
			Telepon: in.Pembeli.Telepon,
			Email: in.Pembeli.Email,
		},
		Produk: models.Catalog{
			NamaProduk: in.Barang.Produk,
			Url: in.Barang.Link,
			Gambar: in.Barang.Gambar,
			Harga: in.Barang.Harga,
			Kategori: in.Barang.Kategori,
		},
		Status: in.Status,
	}
	return res
}

func (s *server) GetCatalog(ctx context.Context, in *v1.Keyword) (*v1.Products, error) {
	keyword := ""
	if in != nil {
		keyword = in.Keyword
	}
	list, err := s.catalog.GetCatalog(keyword)
	fmt.Printf("%+v", list)
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
		Products: arrProducts,
	}
	return result, nil
}

// func (s *server) GetListArticle(ctx context.Context, in *article_grpc.FetchRequest) (*article_grpc.ListArticle, error) {
// 	cursor := ""
// 	num := int64(0)
// 	if in != nil {

// 		cursor = in.Cursor

// 		num = in.Num
// 	}
// 	list, nextCursor, err := s.usecase.Fetch(cursor, num)

// 	if err != nil {
// 		log.Println(err.Error())
// 		return nil, err
// 	}
// 	arrArticle := make([]*article_grpc.Article, len(list))
// 	for i, a := range list {
// 		ar := s.transformArticleRPC(a)
// 		arrArticle[i] = ar
// 	}
// 	result := &article_grpc.ListArticle{
// 		Artilces: arrArticle,
// 		Cursor:   nextCursor,
// 	}
// 	return result, nil
// }

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

// func (s *server) BatchInsert(stream article_grpc.ArticleHandler_BatchInsertServer) error {
// 	errs := make([]*article_grpc.ErrorMessage, 0)
// 	totalSukses := int64(0)
// 	for {
// 		article, err := stream.Recv()
// 		if err == io.EOF {
// 			return stream.SendAndClose(&article_grpc.BatchInsertResponse{
// 				Errors:       errs,
// 				TotalSuccess: totalSukses,
// 			})
// 		}
// 		if err != nil {
// 			log.Println(err.Error())
// 			return err
// 		}
// 		a := s.transformArticleData(article)
// 		res, err := s.usecase.Store(a)
// 		if err != nil {
// 			log.Println(err.Error())
// 			e := &article_grpc.ErrorMessage{
// 				Message: err.Error(),
// 			}
// 			errs = append(errs, e)
// 		}
// 		if res != nil {
// 			totalSukses++
// 		}
// 	}

// }

func (s *server) GetOrders(ctx context.Context, in *v1.UserId) (*v1.Orders,error) {
	listOrder, err := s.order.GetOrders(in.Id)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	arr := make([]*v1.Order, len(listOrder))
	for i, a := range listOrder {
		order := s.transformOrderRPC(&a)
		arr[i] = order
	}

	res := &v1.Orders{
		Orders: arr,
	}
	return res, nil
}

func (s *server) CreateOrder(ctx context.Context, in *v1.Order) (*v1.Response,error) {
	newOrder := s.transformOrderData(in)
	created, err := s.order.CreateOrder(&newOrder)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	res := &v1.Response{
		Success: created.Success,
		Code: created.Code,
		Message: created.Message,
	}
	return res, nil
}

func (s *server) UpdateStatusOrder(ctx context.Context, in *v1.Update) (*v1.Response,error) {
	id := in.Id
	status := in.Status

	updated, err := s.order.UpdateStatusOrder(id, status)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	res := &v1.Response{
		Success: updated.Success,
		Code: updated.Code,
		Message: updated.Message,
	}
	return res, nil
}

func (s *server) DeleteOrder(ctx context.Context, in *v1.OrderId) (*v1.Response,error) {
	deleted, err := s.order.DeleteOrder(in.OrderId)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	res := &v1.Response{
		Success: deleted.Success,
		Code: deleted.Code,
		Message: deleted.Message,
	}
	return res, nil
}