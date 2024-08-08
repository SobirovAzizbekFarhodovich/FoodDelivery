package postgres

import (
	"database/sql"
	pb "delivery/genprotos"
	"errors"
	"fmt"
	"log"
	"strings"
)

type OrderItemsRepo struct {
	db *sql.DB
}

func NewOrderItemsRepo(db *sql.DB) *OrderItemsRepo {
	return &OrderItemsRepo{db: db}
}

func (p *OrderItemsRepo) CreateOrderItems(order *pb.CreateOrderItemsRequest) (*pb.CreateOrderItemsResponse, error) {
	query := `INSERT INTO order_items(order_id,product_id,quantity,options)
	          VALUES($1, $2, $3, $4)
	          RETURNING id, order_id,product_id,quantity,options`
	var res pb.CreateOrderItemsResponse
	err := p.db.QueryRow(query, order.OrderId, order.ProductId, order.Quantity, order.Options).Scan(
		&res.Id,
		&res.OrderId,
		&res.ProductId,
		&res.Quantity,
		&res.Options,
	)
	if err != nil {
		log.Printf("Failed to create order: %v", err)
		return nil, fmt.Errorf("failed to create order: %v", err)
	}

	return &res, nil
}

func (p *OrderItemsRepo) UpdateOrderItems(order *pb.UpdateOrderItemsRequest) (*pb.UpdateOrderItemsResponse, error) {
	query := `UPDATE order_items SET `
	var condition []string
	var args []interface{}

	if order.OrderId != "" && order.OrderId != "string" {
		condition = append(condition, fmt.Sprintf("order_id = $%d", len(args)+1))
		args = append(args, order.OrderId)
	}

	if order.ProductId != "" && order.ProductId != "string" {
		condition = append(condition, fmt.Sprintf("product_id = $%d", len(args)+1))
		args = append(args, order.ProductId)
	}
	if order.Quantity != 0 {
		condition = append(condition, fmt.Sprintf("quantity = $%d", len(args)+1))
		args = append(args, order.Quantity)
	}
	if order.Options != "" && order.Options != "string" {
		condition = append(condition, fmt.Sprintf("options = $%d", len(args)+1))
		args = append(args, order.Options)
	}
	if len(condition) == 0 {
		return nil, errors.New("nothing to update")
	}

	query += strings.Join(condition, ", ")
	query += fmt.Sprintf(" WHERE id = $%d and deleted_at = 0 RETURNING id,order_id,quantity,options,product_id", len(args)+1)
	args = append(args, order.Id)

	res := pb.UpdateOrderItemsResponse{}
	row := p.db.QueryRow(query, args...)
	err := row.Scan(&res.Id, &res.OrderId, &res.Quantity, &res.Options, &res.ProductId)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (p *OrderItemsRepo) DeleteOrderItems(id *pb.DeleteOrderItemsRequest) (*pb.DeleteOrderItemsResponse, error) {
	query := `UPDATE order_items SET deleted_at = EXTRACT(EPOCH FROM NOW()) WHERE id = $1 and deleted_at = 0`
	res, err := p.db.Exec(query, id.Id)
	if err != nil {
		log.Println("Error while delete order", err)
		return nil, err
	}
	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("order with id %s not found", id.Id)
	}
	log.Println("Successfully deleted order items")

	return &pb.DeleteOrderItemsResponse{}, nil
}

func (p *OrderItemsRepo) GetByIdOrderItems(id *pb.GetByIdOrderItemsRequest) (*pb.GetByIdOrderItemsResponse, error) {
	query := `SELECT id,order_id,product_id,quantity,options FROM order_items WHERE id = $1 and deleted_at = 0`
	row := p.db.QueryRow(query, id.Id)
	res := pb.GetByIdOrderItemsResponse{}
	err := row.Scan(&res.Id, &res.OrderId, &res.ProductId, &res.Quantity, &res.Options)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("order not found")
		}
		return nil, err
	}
	return &res, nil
}

func (p *OrderItemsRepo) GetAllOrderItems(flt *pb.GetAllOrderItemsRequest) (*pb.GetAllOrderItemsResponse, error) {
	limit := int32(10)
	offset := int32(0)
	if flt.Limit > 0{
		limit = flt.Limit
	}
	if flt.Page > 0{
		offset = (flt.Page - 1) * limit
	} 
	query := `
		SELECT id,order_id,product_id,quantity,options
		FROM order_items
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`
	rows, err := p.db.Query(query, limit, offset)
	if err != nil {
		log.Printf("Failed to get all orders: %v", err)
		return nil, fmt.Errorf("failed to get all orders: %v", err)
	}
	defer rows.Close()

	var orders []*pb.GetByIdOrderItemsResponse
	for rows.Next() {
		var res pb.GetByIdOrderItemsResponse
		if err := rows.Scan(
			&res.Id,
			&res.OrderId,
			&res.ProductId,
			&res.Quantity,
			&res.Options,
		); err != nil {
			log.Printf("Failed to scan order row: %v", err)
			return nil, fmt.Errorf("failed to scan order row: %v", err)
		}
		orders = append(orders, &res)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Failed to iterate over orders: %v", err)
		return nil, fmt.Errorf("failed to iterate over orders: %v", err)
	}

	return &pb.GetAllOrderItemsResponse{
		Orders: orders,
	}, nil
}
