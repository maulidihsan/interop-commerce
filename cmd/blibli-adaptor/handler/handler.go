package handler

import (
	"log"
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"github.com/maulidihsan/flashdeal-webservice/pkg/catalog/usecase"
	"github.com/maulidihsan/flashdeal-webservice/pkg/v1"
	"github.com/maulidihsan/flashdeal-webservice/pkg/models"
)

func NewCatalogServerGrpc(gserver *grpc.Server, catalogUsecase usecase.CatalogUsecase) {

	catalogServer := &server{
		usecase: catalogUsecase,
	}

	v1.RegisterCatalogServiceServer(gserver, catalogServer)
	reflection.Register(gserver)
}

type server struct {
	usecase usecase.CatalogUsecase
}

func (s *server) transformArticleRPC(ar *models.Catalog) *v1.Product {
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

func (s *server) GetCatalog(ctx context.Context, in *v1.Keyword) (*v1.Products, error) {
	keyword := ""
	if in != nil {
		keyword = in.Keyword
	}
	list, err := s.usecase.GetCatalog(keyword)
	fmt.Printf("%+v", list)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	arrProducts := make([]*v1.Product, len(list))
	for i, a := range list {
		ar := s.transformArticleRPC(&a)
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
	update, err := s.usecase.UpdateCatalog(arrProducts)
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