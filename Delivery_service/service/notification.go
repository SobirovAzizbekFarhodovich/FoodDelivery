package service

import (
	"context"
	pb "delivery/genprotos"
	"delivery/storage"
)

type NotificationService struct {
	storage storage.StorageI
	pb.UnimplementedNotificationServiceServer
}

func NewNotificationService(storage storage.StorageI) *NotificationService {
	return &NotificationService{storage: storage}
}

func (c *NotificationService) CreateNotification(ctx context.Context, not *pb.CreateNotificationRequest) (*pb.CreateNotificationResponse, error) {
	_, err := c.storage.Notification().CreateNotification(not)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (c *NotificationService) GetNotification(ctx context.Context, flt *pb.GetNotificationsRequest) (*pb.GetNotificationsResponse, error) {
	res, err := c.storage.Notification().GetNotification(flt)
	if err != nil {
		return nil, err
	}
	return res, nil
}
