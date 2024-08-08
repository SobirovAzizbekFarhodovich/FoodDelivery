package handler

import (
	pb "api/genprotos/delivery"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateTask godoc
// @Summary Create a new task
// @Description Create a new task
// @Tags task
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param task body pb.CreateTaskRequest true "Create task"
// @Success 201 {object} pb.CreateTaskResponse
// @Failure 400 {string} string "Bad Request"
// @Router /task/create [post]
func (h *Handler) CreateTask(ctx *gin.Context) {
	req := &pb.CreateTaskRequest{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	res, err := h.Task.CreateTask(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, res)
}

// UpdateTask godoc
// @Summary Update an existing task
// @Description Update an existing task
// @Tags task
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param task body pb.UpdateTaskRequest true "Update task"
// @Success 200 {object} pb.UpdateTaskResponse
// @Failure 400 {string} string "Bad Request"
// @Router /task/update [put]
func (h *Handler) UpdateTask(ctx *gin.Context) {
	req := &pb.UpdateTaskRequest{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	res, err := h.Task.UpdateTask(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// GetByIdTask godoc
// @Summary Get a task by ID
// @Description Get a task by ID
// @Tags task
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Task ID"
// @Success 200 {object} pb.GetByIdTaskResponse
// @Failure 400 {string} string "Bad Request"
// @Router /task/get/{id} [get]
func (h *Handler) GetByIdTask(ctx *gin.Context) {
	req := &pb.GetByIdTaskRequest{
		Id: ctx.Param("id"),
	}
	res, err := h.Task.GetByIdTask(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// GetAllTasks godoc
// @Summary Get all tasks
// @Description Get all tasks
// @Tags task
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Success 200 {object} pb.GetAllTasksResponse
// @Failure 400 {string} string "Bad Request"
// @Router /task/get [get]
func (h *Handler) GetAllTasks(ctx *gin.Context) {
	defaultLimit := 10
	defaultPage := 1

	limitStr := ctx.Query("limit")
	pageStr := ctx.Query("page")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = defaultLimit
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = defaultPage
	}

	req := &pb.GetAllTasksRequest{
		Limit: int32(limit),
		Page:  int32(page),
	}

	res, err := h.Task.GetAllTasks(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// DeleteTask godoc
// @Summary Delete a task
// @Description Delete a task
// @Tags task
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Task ID"
// @Success 200 {object} string "OK"
// @Failure 400 {string} string "Bad Request"
// @Router /task/delete/{id} [delete]
func (h *Handler) DeleteTask(ctx *gin.Context) {
	req := &pb.DeleteTaskRequest{
		Id: ctx.Param("id"),
	}
	_, err := h.Task.DeleteTask(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message":"Task deleted successfully"})
}
