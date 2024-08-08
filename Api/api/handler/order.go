package handler

import (
	pb "api/genprotos/delivery"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateOrder godoc
// @Summary Create a new order
// @Description Create a new order
// @Tags order
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param order body pb.CreateOrderRequest true "Create order"
// @Success 201 {object} pb.CreateOrderResponse
// @Failure 400 {string} string "Bad Request"
// @Router /order/create [post]
func (h *Handler) CreateOrder(ctx *gin.Context) {
	req := &pb.CreateOrderRequest{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	res, err := h.Order.CreateOrder(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, res)
}

// UpdateOrder godoc
// @Summary Update an existing order
// @Description Update an existing order
// @Tags order
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param order body pb.UpdateOrderRequest true "Update order"
// @Success 200 {object} pb.UpdateOrderResponse
// @Failure 400 {string} string "Bad Request"
// @Router /order/update [put]
func (h *Handler) UpdateOrder(ctx *gin.Context) {
	req := &pb.UpdateOrderRequest{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	res, err := h.Order.UpdateOrder(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// DeleteOrder godoc
// @Summary Delete an order
// @Description Delete an order
// @Tags order
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Order ID"
// @Success 200 {object} string "OK"
// @Failure 400 {string} string "Bad Request"
// @Router /order/delete/{id} [delete]
func (h *Handler) DeleteOrder(ctx *gin.Context) {
	req := &pb.DeleteOrderRequest{
		Id: ctx.Param("id"),
	}
	_, err := h.Order.DeleteOrder(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}

// GetByIdOrder godoc
// @Summary Get an order by ID
// @Description Get an order by ID
// @Tags order
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Order ID"
// @Success 200 {object} pb.GetByIdOrderResponse
// @Failure 400 {string} string "Bad Request"
// @Router /order/get/{id} [get]
func (h *Handler) GetByIdOrder(ctx *gin.Context) {
	req := &pb.GetByIdOrderRequest{
		Id: ctx.Param("id"),
	}
	res, err := h.Order.GetByIdOrder(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// GetAllOrders godoc
// @Summary Get all orders
// @Description Get all orders
// @Tags order
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Success 200 {object} pb.GetAllOrdersResponse
// @Failure 400 {string} string "Bad Request"
// @Router /order/get [get]
func (h *Handler) GetAllOrders(ctx *gin.Context) {
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

	req := &pb.GetAllOrdersRequest{
		Limit: int32(limit),
		Page:  int32(page),
	}

	res, err := h.Order.GetAllOrders(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, res)
}


// GetByLocation godoc
// @Summary Get orders by location
// @Description Retrieve orders filtered by a specific location
// @Tags order
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param location path string true "Order Location"
// @Success 200 {object} pb.GetByLocationResponse
// @Failure 400 {string} string "Bad Request"
// @Router /order/{location} [get]
func (h *Handler) GetByLocation(ctx *gin.Context) {
	req := pb.GetByLocationRequest{
		Location: ctx.Param("location"),
	}
	res, err := h.Order.GetByLocation(context.Background(), &req)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	ctx.JSON(200, res)
}


// UpdateStatus godoc
// @Summary Update the status of an order
// @Description Modify the status of an order by its ID
// @Tags order
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param status body pb.UpdateStatusRequest true "Update Status"
// @Success 200 {object} string "Update status successfully"
// @Failure 400 {string} string "Bad Request"
// @Router /order/status [put]
func (h *Handler) UpdateStatus(ctx *gin.Context){
	req := &pb.UpdateStatusRequest{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	_, err := h.Order.UpdateStatus(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message":"Update status successfully"})
}

// GetHistory godoc
// @Summary Get the order history
// @Description Retrieve a paginated list of delivered orders
// @Tags order
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Success 200 {object} pb.GetHistoryResponse
// @Failure 400 {string} string "Bad Request"
// @Router /order/history [get]
func(h *Handler) GetHistory(ctx *gin.Context){
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

	req := &pb.GetHistoryRequest{
		Limit: int32(limit),
		Page:  int32(page),
	}

	res, err := h.Order.GetHistory(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, res)
}