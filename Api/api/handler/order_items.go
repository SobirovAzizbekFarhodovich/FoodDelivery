package handler

import (
	pb "api/genprotos/delivery"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateOrderItems godoc
// @Summary Create a new order item
// @Description Create a new order item
// @Tags order_items
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param order_items body pb.CreateOrderItemsRequest true "Create order item"
// @Success 201 {object} pb.CreateOrderItemsResponse
// @Failure 400 {string} string "Bad Request"
// @Router /order_items/create [post]
func (h *Handler) CreateOrderItems(ctx *gin.Context) {
	req := &pb.CreateOrderItemsRequest{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	res, err := h.OrderItems.CreateOrderItems(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, res)
}

// UpdateOrderItems godoc
// @Summary Update an existing order item
// @Description Update an existing order item
// @Tags order_items
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param order_items body pb.UpdateOrderItemsRequest true "Update order item"
// @Success 200 {object} pb.UpdateOrderItemsResponse
// @Failure 400 {string} string "Bad Request"
// @Router /order_items/update [put]
func (h *Handler) UpdateOrderItems(ctx *gin.Context) {
	req := &pb.UpdateOrderItemsRequest{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	res, err := h.OrderItems.UpdateOrderItems(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// DeleteOrderItems godoc
// @Summary Delete an order item
// @Description Delete an order item
// @Tags order_items
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Order Item ID"
// @Success 200 {object} string "OK"
// @Failure 400 {string} string "Bad Request"
// @Router /order_items/delete/{id} [delete]
func (h *Handler) DeleteOrderItems(ctx *gin.Context) {
	req := &pb.DeleteOrderItemsRequest{
		Id: ctx.Param("id"),
	}
	_, err := h.OrderItems.DeleteOrderItems(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message":"Order item deleted successfully"})
}

// GetByIdOrderItems godoc
// @Summary Get an order item by ID
// @Description Get an order item by ID
// @Tags order_items
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Order Item ID"
// @Success 200 {object} pb.GetByIdOrderItemsResponse
// @Failure 400 {string} string "Bad Request"
// @Router /order_items/get/{id} [get]
func (h *Handler) GetByIdOrderItems(ctx *gin.Context) {
	req := &pb.GetByIdOrderItemsRequest{
		Id: ctx.Param("id"),
	}
	res, err := h.OrderItems.GetByIdOrderItems(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// GetAllOrderItems godoc
// @Summary Get all order items
// @Description Get all order items
// @Tags order_items
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Success 200 {object} pb.GetAllOrderItemsResponse
// @Failure 400 {string} string "Bad Request"
// @Router /order_items/get [get]
func (h *Handler) GetAllOrderItems(ctx *gin.Context) {
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

	req := &pb.GetAllOrderItemsRequest{
		Limit: int32(limit),
		Page:  int32(page),
	}

	res, err := h.OrderItems.GetAllOrderItems(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, res)
}
