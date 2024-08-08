package service

import (
	"context"

	pb "delivery/genprotos"
	"delivery/storage"
)

type TaskService struct {
	storage storage.StorageI
	pb.UnimplementedTaskServiceServer
}

func NewTaskService(storage storage.StorageI) *TaskService {
	return &TaskService{storage: storage}
}

func (c *TaskService) CreateTask(ctx context.Context, task *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
	res, err := c.storage.Task().CreateTask(task)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *TaskService) UpdateTask(ctx context.Context, task *pb.UpdateTaskRequest) (*pb.UpdateTaskResponse, error) {
	res, err := c.storage.Task().UpdateTask(task)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *TaskService) DeleteTask(ctx context.Context, id *pb.DeleteTaskRequest) (*pb.DeleteTaskResponse, error) {
	_, err := c.storage.Task().DeleteTask(id)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (c *TaskService) GetByIdTask(ctx context.Context, id *pb.GetByIdTaskRequest) (*pb.GetByIdTaskResponse, error) {
	res, err := c.storage.Task().GetByIdTask(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *TaskService) GetAllTask(ctx context.Context, flt *pb.GetAllTasksRequest) (*pb.GetAllTasksResponse, error) {
	res, err := c.storage.Task().GetAllTask(flt)
	if err != nil {
		return nil, err
	}
	return res, nil
}
