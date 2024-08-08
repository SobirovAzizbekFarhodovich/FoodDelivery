package main

import (
	"log"
	"net"

	pb "delivery/genprotos"
	"delivery/service"
	"delivery/storage/postgres"
	"google.golang.org/grpc"
)

func main(){
	db, err := postgres.ConnectDb()
	if err != nil {
		log.Fatal(err)
	}

	liss, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterCartServiceServer(s, service.NewCartService(db))
	pb.RegisterOrderItemsServiceServer(s, service.NewOrderItemsService(db))
	pb.RegisterOrderServiceServer(s, service.NewOrderService(db))
	pb.RegisterProductServiceServer(s, service.NewProductService(db))
	pb.RegisterTaskServiceServer(s,service.NewTaskService(db))
	pb.RegisterNotificationServiceServer(s, service.NewNotificationService(db))

	log.Printf("server listening at %v", liss.Addr())
	if err := s.Serve(liss); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
