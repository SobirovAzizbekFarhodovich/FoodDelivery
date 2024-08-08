package service

import (
	"context"

	pb "delivery/genprotos"
	"delivery/storage"
)

type ProductService struct{
	storage storage.StorageI
	pb.UnimplementedProductServiceServer
}

func NewProductService (storage storage.StorageI) *ProductService{
	return &ProductService{storage: storage}
}

func (c *ProductService) CreateProduct(ctx context.Context,product *pb.CreateProductRequest) (*pb.CreateProductResponse, error){
	res, err := c.storage.Product().CreateProduct(product)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *ProductService) DeleteProduct(ctx context.Context,id *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error){
	_, err := c.storage.Product().DeleteProduct(id)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (c *ProductService) UpdateProduct(ctx context.Context,product *pb.UpdateProductRequest) (*pb.UpdateProductResponse, error){
	res, err := c.storage.Product().UpdateProduct(product)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *ProductService) GetByIdProduct(ctx context.Context,id *pb.GetByIdProductRequest) (*pb.GetByIdProductResponse, error){
	res, err := c.storage.Product().GetByIdProduct(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *ProductService) GetAllProduct(ctx context.Context,flt *pb.GetAllProductsRequest) (*pb.GetAllProductsResponse, error){
	res, err := c.storage.Product().GetAllProduct(flt)
	if err != nil {
		return nil, err
	}
	return res, nil
}