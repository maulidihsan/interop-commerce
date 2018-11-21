package handler

import (
	"context"
	"log"
	
	"github.com/maulidihsan/flashdeal-webservice/pkg/v1"
	"github.com/maulidihsan/flashdeal-webservice/pkg/models"
)

func (s *server) transformOrderRPC(in *models.Order) *v1.Order {
	if in == nil {
		return nil
	}
	res := &v1.Order{
		Vendor: v1.Vendor_BLIBLI,
		Id: in.Id,
		Pembeli: &v1.Person{
			Nama: in.Pembeli.Nama,
			Alamat: in.Pembeli.Alamat,
			Telepon: in.Pembeli.Telepon,
			Email: in.Pembeli.Email,
		},
		Barang: &v1.Product{
			Id: in.Produk.Id,
			Produk: in.Produk.NamaProduk,
			Link: in.Produk.Url,
			Gambar: in.Produk.Gambar,
			Harga: in.Produk.Harga,
			Kategori: in.Produk.Kategori,
		},
		Kuantitas: in.Kuantitas,
		Status: in.Status,
		CreatedAt: in.CreatedAt.Unix(),
	}
	return res
}

func (s *server) transformOrderData(in *v1.Order) (*models.Order, error) {
	product, err := s.catalog.GetById(in.Barang.GetId())
	if (err != nil) {
		return nil, err
	}
	res := &models.Order{
		Pembeli: &models.User{
			Nama: in.Pembeli.Nama,
			Alamat: in.Pembeli.Alamat,
			Telepon: in.Pembeli.Telepon,
			Email: in.Pembeli.Email,
		},
		Produk: models.Catalog{
			Id: in.Barang.Id,
			NamaProduk: product.NamaProduk,
			Url: product.Url,
			Gambar: product.Gambar,
			Harga: product.Harga,
			Kategori: product.Kategori,
		},
		Kuantitas: in.Kuantitas,
		Status: "diterima",
	}
	return res, nil
}

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
	if(in.Vendor.String() == "BLIBLI") {
		newOrder, err := s.transformOrderData(in)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		created, err := s.order.CreateOrder(newOrder)
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
	} else {
		res := &v1.Response{
			Success: false,
			Code: 400,
			Message: "Wrong Service Destination",
		}
		return res, nil
	}
}

func (s *server) UpdateStatusOrder(ctx context.Context, in *v1.Update) (*v1.Response,error) {
	if(in.Vendor == 0) {
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
	} else {
		res := &v1.Response{
			Success: false,
			Code: 400,
			Message: "Wrong Service Destination",
		}
		return res, nil
	}
}

func (s *server) DeleteOrder(ctx context.Context, in *v1.OrderId) (*v1.Response,error) {
	if(in.Vendor == 0) {
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
	} else {
		res := &v1.Response{
			Success: false,
			Code: 400,
			Message: "Wrong Service Destination",
		}
		return res, nil
	}
}