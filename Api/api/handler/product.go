package handler

import (
	pb "api/genprotos/delivery"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product
// @Tags product
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param product body pb.CreateProductRequest true "Create product"
// @Success 201 {object} pb.CreateProductResponse
// @Failure 400 {string} string "Bad Request"
// @Router /product/create [post]
func (h *Handler) CreateProduct(ctx *gin.Context) {
	req := &pb.CreateProductRequest{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	res, err := h.Product.CreateProduct(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, res)
}

// UpdateProduct godoc
// @Summary Update an existing product
// @Description Update an existing product
// @Tags product
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param product body pb.UpdateProductRequest true "Update product"
// @Success 200 {object} pb.UpdateProductResponse
// @Failure 400 {string} string "Bad Request"
// @Router /product/update [put]
func (h *Handler) UpdateProduct(ctx *gin.Context) {
	req := &pb.UpdateProductRequest{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	res, err := h.Product.UpdateProduct(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// DeleteProduct godoc
// @Summary Delete a product
// @Description Delete a product
// @Tags product
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Product ID"
// @Success 200 {object} string "OK"
// @Failure 400 {string} string "Bad Request"
// @Router /product/delete/{id} [delete]
func (h *Handler) DeleteProduct(ctx *gin.Context) {
	req := &pb.DeleteProductRequest{
		Id: ctx.Param("id"),
	}
	_, err := h.Product.DeleteProduct(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message":"Product deleted successfully"})
}

// GetByIdProduct godoc
// @Summary Get a product by ID
// @Description Get a product by ID
// @Tags product
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Product ID"
// @Success 200 {object} pb.GetByIdProductResponse
// @Failure 400 {string} string "Bad Request"
// @Router /product/get/{id} [get]
func (h *Handler) GetByIdProduct(ctx *gin.Context) {
	req := &pb.GetByIdProductRequest{
		Id: ctx.Param("id"),
	}
	res, err := h.Product.GetByIdProduct(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// GetAllProducts godoc
// @Summary Get all products
// @Description Get all products
// @Tags product
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Success 200 {object} pb.GetAllProductsResponse
// @Failure 400 {string} string "Bad Request"
// @Router /product/get [get]
func (h *Handler) GetAllProducts(ctx *gin.Context) {
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

	req := &pb.GetAllProductsRequest{
		Limit: int32(limit),
		Page:  int32(page),
	}

	res, err := h.Product.GetAllProducts(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, res)
}
