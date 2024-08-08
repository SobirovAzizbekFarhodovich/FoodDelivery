package handler

import (
	delivery "api/genprotos/delivery"

	"google.golang.org/grpc"
)

type Handler struct {
	Cart       delivery.CartServiceClient
	OrderItems delivery.OrderItemsServiceClient
	Order      delivery.OrderServiceClient
	Product    delivery.ProductServiceClient
	Task       delivery.TaskServiceClient
}

func NewHandler(deliveryConn *grpc.ClientConn) *Handler{
	return &Handler{
		Cart: delivery.NewCartServiceClient(deliveryConn),
		OrderItems: delivery.NewOrderItemsServiceClient(deliveryConn),
		Order: delivery.NewOrderServiceClient(deliveryConn),
		Product: delivery.NewProductServiceClient(deliveryConn),
		Task: delivery.NewTaskServiceClient(deliveryConn),
	}
}
