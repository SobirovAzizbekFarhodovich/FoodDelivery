package storage

import (
	pb "delivery/genprotos"
)

type StorageI interface {
	Cart() CartI
	Order() OrderI
	OrderItems() OrderItemsI
	Product() ProductI
	Task() TaskI
	Notification() NotificationI
}

type CartI interface {
	CreateCart(cart *pb.CreateCartRequest) (*pb.CreateCartResponse, error)
	UpdateCart(order *pb.UpdateCartRequest) (*pb.UpdateCartResponse, error)
	DeleteCart(id *pb.DeleteCartRequest) (*pb.DeleteCartResponse, error)
	GetByIdCart(id *pb.GetByIdCartRequest) (*pb.GetByIdCartResponse, error)
	GetAllCart(flt *pb.GetAllCartRequest) (*pb.GetAllCartResponse, error)
}

type OrderI interface {
	CreateOrder(order *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error)
	UpdateOrder(order *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error)
	DeleteOrder(id *pb.DeleteOrderRequest) (*pb.DeleteOrderResponse, error)
	GetByIdOrder(id *pb.GetByIdOrderRequest) (*pb.GetByIdOrderResponse, error)
	GetAllOrder(flt *pb.GetAllOrdersRequest) (*pb.GetAllOrdersResponse, error)
	GetByLocation(order *pb.GetByLocationRequest) (*pb.GetByLocationResponse, error)
	UpdateStatus(status *pb.UpdateStatusRequest) (*pb.UpdateStatusResponse, error)
	GetHistory(history *pb.GetHistoryRequest) (*pb.GetHistoryResponse, error)
}

type OrderItemsI interface {
	CreateOrderItems(order *pb.CreateOrderItemsRequest) (*pb.CreateOrderItemsResponse, error)
	UpdateOrderItems(order *pb.UpdateOrderItemsRequest) (*pb.UpdateOrderItemsResponse, error)
	DeleteOrderItems(id *pb.DeleteOrderItemsRequest) (*pb.DeleteOrderItemsResponse, error)
	GetByIdOrderItems(id *pb.GetByIdOrderItemsRequest) (*pb.GetByIdOrderItemsResponse, error)
	GetAllOrderItems(flt *pb.GetAllOrderItemsRequest) (*pb.GetAllOrderItemsResponse, error)
}

type ProductI interface {
	CreateProduct(product *pb.CreateProductRequest) (*pb.CreateProductResponse, error)
	UpdateProduct(product *pb.UpdateProductRequest) (*pb.UpdateProductResponse, error)
	DeleteProduct(id *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error)
	GetByIdProduct(id *pb.GetByIdProductRequest) (*pb.GetByIdProductResponse, error)
	GetAllProduct(flt *pb.GetAllProductsRequest) (*pb.GetAllProductsResponse, error)
}

type TaskI interface {
	CreateTask(task *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error)
	UpdateTask(task *pb.UpdateTaskRequest) (*pb.UpdateTaskResponse, error)
	DeleteTask(id *pb.DeleteTaskRequest) (*pb.DeleteTaskResponse, error)
	GetByIdTask(id *pb.GetByIdTaskRequest) (*pb.GetByIdTaskResponse, error)
	GetAllTask(flt *pb.GetAllTasksRequest) (*pb.GetAllTasksResponse, error)
}

type NotificationI interface {
	CreateNotification(not *pb.CreateNotificationRequest) (*pb.CreateNotificationResponse, error)
	GetNotification(flt *pb.GetNotificationsRequest) (*pb.GetNotificationsResponse, error)
}
