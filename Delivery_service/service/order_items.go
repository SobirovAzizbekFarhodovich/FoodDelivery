package service

import (
	"context"

	pb "delivery/genprotos"
	"delivery/storage"
)

type OrderItemsService struct{
	storage storage.StorageI
	pb.UnimplementedOrderItemsServiceServer
}

func NewOrderItemsService (storage storage.StorageI) *OrderItemsService{
	return &OrderItemsService{storage: storage}
}

func (c *OrderItemsService) CreateOrderItems(ctx context.Context,order *pb.CreateOrderItemsRequest) (*pb.CreateOrderItemsResponse, error){
	res, err := c.storage.OrderItems().CreateOrderItems(order)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *OrderItemsService) UpdateOrderItems(ctx context.Context,order *pb.UpdateOrderItemsRequest) (*pb.UpdateOrderItemsResponse, error){
	res, err := c.storage.OrderItems().UpdateOrderItems(order)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *OrderItemsService) DeleteOrderItems(ctx context.Context,id *pb.DeleteOrderItemsRequest) (*pb.DeleteOrderItemsResponse, error){
	_, err := c.storage.OrderItems().DeleteOrderItems(id)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (c *OrderItemsService) GetByIdOrderItems(ctx context.Context,id *pb.GetByIdOrderItemsRequest) (*pb.GetByIdOrderItemsResponse, error){
	res, err := c.storage.OrderItems().GetByIdOrderItems(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *OrderItemsService) GetAllOrderItems(ctx context.Context,flt *pb.GetAllOrderItemsRequest) (*pb.GetAllOrderItemsResponse, error){
	res, err := c.storage.OrderItems().GetAllOrderItems(flt)
	if err != nil {
		return nil, err
	}
	return res, nil
}