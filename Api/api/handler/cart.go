package handler

import (
	pb "api/genprotos/delivery"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateCart godoc
// @Summary Create a new cart
// @Description Create a new cart
// @Tags cart
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param cart body pb.CreateCartRequest true "Create cart"
// @Success 201 {object} pb.CreateCartResponse
// @Failure 400 {string} string "Bad Request"
// @Router /cart/create [post]
func (h *Handler) CreateCart(ctx *gin.Context) {
    req := &pb.CreateCartRequest{}
    if err := ctx.BindJSON(req); err != nil {
        ctx.JSON(http.StatusBadRequest, err.Error())
        return
    }
    res, err := h.Cart.CreateCart(context.Background(), req)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, err.Error())
        return
    }
    ctx.JSON(http.StatusCreated, res)
}

// UpdateCart godoc
// @Summary Update an existing cart
// @Description Update an existing cart
// @Tags cart
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param cart body pb.UpdateCartRequest true "Update cart"
// @Success 200 {object} pb.UpdateCartResponse
// @Failure 400 {string} string "Bad Request"
// @Router /cart/update [put]
func (h *Handler) UpdateCart(ctx *gin.Context) {
	req := &pb.UpdateCartRequest{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	res, err := h.Cart.UpdateCart(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// DeleteCart godoc
// @Summary Delete a cart
// @Description Delete a cart
// @Tags cart
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Cart ID"
// @Success 200 {object} string "OK"
// @Failure 400 {string} string "Bad Request"
// @Router /cart/delete/{id} [delete]
func (h *Handler) DeleteCart(ctx *gin.Context) {
	req := &pb.DeleteCartRequest{
		Id: ctx.Param("id"),
	}
	_, err := h.Cart.DeleteCart(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message":"Cart deleted succussfully"})
}

// GetByIdCart godoc
// @Summary Get a cart by ID
// @Description Get a cart by ID
// @Tags cart
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Cart ID"
// @Success 200 {object} pb.GetByIdCartResponse
// @Failure 400 {string} string "Bad Request"
// @Router /cart/get/{id} [get]
func (h *Handler) GetByIdCart(ctx *gin.Context) {
	req := &pb.GetByIdCartRequest{
		Id: ctx.Param("id"),
	}
	res, err := h.Cart.GetByIdCart(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// GetAllCart godoc
// @Summary Get all carts
// @Description Get all carts
// @Tags cart
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Success 200 {object} pb.GetAllCartResponse
// @Failure 400 {string} string "Bad Request"
// @Router /cart/get [get]
func (h *Handler) GetAllCart(ctx *gin.Context) {
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

	req := &pb.GetAllCartRequest{
		Limit: int32(limit),
		Page:  int32(page),
	}

	res, err := h.Cart.GetAllCart(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, res)
}
