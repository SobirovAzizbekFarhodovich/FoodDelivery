package service

import (
	"context"

	pb "delivery/genprotos"
	"delivery/storage"
)

type CartService struct{
	storage storage.StorageI
	pb.UnimplementedCartServiceServer
}

func NewCartService (storage storage.StorageI) *CartService{
	return &CartService{storage: storage}
}

func (c *CartService) CreateCart(ctx context.Context,cart *pb.CreateCartRequest) (*pb.CreateCartResponse, error){
	res, err := c.storage.Cart().CreateCart(cart)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *CartService) UpdateCart(ctx context.Context,order *pb.UpdateCartRequest) (*pb.UpdateCartResponse, error){
	res, err := c.storage.Cart().UpdateCart(order)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *CartService) DeleteCart(ctx context.Context,id *pb.DeleteCartRequest) (*pb.DeleteCartResponse, error){
	_, err := c.storage.Cart().DeleteCart(id)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (c *CartService) GetByIdCart(ctx context.Context,id *pb.GetByIdCartRequest) (*pb.GetByIdCartResponse, error){
	res, err := c.storage.Cart().GetByIdCart(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *CartService) GetAllCart(ctx context.Context,flt *pb.GetAllCartRequest) (*pb.GetAllCartResponse, error){
	res, err := c.storage.Cart().GetAllCart(flt)
	if err != nil {
		return nil, err
	}
	return res, nil
}