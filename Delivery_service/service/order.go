package service

import (
	"context"

	pb "delivery/genprotos"
	"delivery/storage"
)

type OrderService struct{
	storage storage.StorageI
	pb.UnimplementedOrderServiceServer
}

func NewOrderService (storage storage.StorageI) *OrderService{
	return &OrderService{storage: storage}
}

func (c *OrderService) CreateOrder(ctx context.Context,order *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error){
	res, err := c.storage.Order().CreateOrder(order)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *OrderService) DeleteOrder(ctx context.Context,id *pb.DeleteOrderRequest) (*pb.DeleteOrderResponse, error){
	_, err := c.storage.Order().DeleteOrder(id)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (c *OrderService) UpdateOrder(ctx context.Context,order *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error){
	res, err := c.storage.Order().UpdateOrder(order)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *OrderService) GetByIdOrder(ctx context.Context,id *pb.GetByIdOrderRequest) (*pb.GetByIdOrderResponse, error){
	res, err := c.storage.Order().GetByIdOrder(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *OrderService) GetAllOrder(ctx context.Context,flt *pb.GetAllOrdersRequest) (*pb.GetAllOrdersResponse, error){
	res, err := c.storage.Order().GetAllOrder(flt)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *OrderService) GetByLocation(ctx context.Context,location *pb.GetByLocationRequest)(*pb.GetByLocationResponse, error){
	res, err := c.storage.Order().GetByLocation(location)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *OrderService) UpdateStatus(ctx context.Context, status *pb.UpdateStatusRequest)(*pb.UpdateStatusResponse, error){
	_, err := c.storage.Order().UpdateStatus(status)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (c *OrderService) GetHistory(ctx context.Context, history *pb.GetHistoryRequest)(*pb.GetHistoryResponse, error){
	res, err := c.storage.Order().GetHistory(history)
	if err != nil {
		return nil, err
	}
	return res, nil
}