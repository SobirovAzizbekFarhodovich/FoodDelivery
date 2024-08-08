package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	pb "delivery/genprotos"
)

type OrderRepo struct {
	db *sql.DB
}

func NewOrderRepo(db *sql.DB) *OrderRepo {
	return &OrderRepo{db: db}
}

func (p *OrderRepo) CreateOrder(order *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	query := `INSERT INTO orders(user_id,total_amount,status,courier_id,location,delivery_schedule)
	          VALUES($1, $2, $3, $4, $5, $6)
	          RETURNING id, user_id,total_amount,status,courier_id,location,delivery_schedule`
	var res pb.CreateOrderResponse
	err := p.db.QueryRow(query, order.UserId, order.TotalAmount, order.Status, order.CourierId, order.Location, order.DeliverySchedule).Scan(
		&res.Id,
		&res.UserId,
		&res.TotalAmount,
		&res.Status,
		&res.CourierId,
		&res.Location,
		&res.DeliverySchedule,
	)
	if err != nil {
		log.Printf("Failed to create order: %v", err)
		return nil, fmt.Errorf("failed to create order: %v", err)
	}

	return &res, nil
}

func (p *OrderRepo) UpdateOrder(order *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error) {
	query := `UPDATE orders SET `
	var condition []string
	var args []interface{}

	if order.Status != "" && order.Status != "string" {
		condition = append(condition, fmt.Sprintf("status = $%d", len(args)+1))
		args = append(args, order.Status)
	}

	if order.CourierId != "" && order.CourierId != "string" {
		condition = append(condition, fmt.Sprintf("courier_id = $%d", len(args)+1))
		args = append(args, order.CourierId)
	}
	if order.TotalAmount != 0 {
		condition = append(condition, fmt.Sprintf("total_amount = $%d", len(args)+1))
		args = append(args, order.TotalAmount)
	}
	if order.Location != "" && order.Location != "string" {
		condition = append(condition, fmt.Sprintf("location = $%d", len(args)+1))
		args = append(args, order.Location)
	}
	if order.DeliverySchedule != "" && order.DeliverySchedule != "string" {
		condition = append(condition, fmt.Sprintf("delivery_schedule = $%d", len(args)+1))
		args = append(args, order.DeliverySchedule)
	}
	if len(condition) == 0 {
		return nil, errors.New("nothing to update")
	}

	query += strings.Join(condition, ", ")
	query += fmt.Sprintf(" WHERE id = $%d and deleted_at = 0 RETURNING id,user_id,total_amount,status,courier_id,location,delivery_schedule", len(args)+1)
	args = append(args, order.Id)

	res := pb.UpdateOrderResponse{}
	row := p.db.QueryRow(query, args...)
	err := row.Scan(&res.Id, &res.UserId, &res.TotalAmount, &res.Status, &res.CourierId, &res.Location, &res.DeliverySchedule)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (p *OrderRepo) DeleteOrder(id *pb.DeleteOrderRequest) (*pb.DeleteOrderResponse, error) {
	query := `UPDATE orders SET deleted_at = EXTRACT(EPOCH FROM NOW()) WHERE id = $1 and deleted_at = 0`
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
	log.Println("Successfully deleted order")

	return &pb.DeleteOrderResponse{}, nil
}

func (p *OrderRepo) GetByIdOrder(id *pb.GetByIdOrderRequest) (*pb.GetByIdOrderResponse, error) {
	query := `SELECT id,user_id,total_amount,status,courier_id,location,delivery_schedule FROM orders WHERE id = $1 and deleted_at = 0`
	row := p.db.QueryRow(query, id.Id)
	res := pb.GetByIdOrderResponse{}
	err := row.Scan(&res.Id, &res.UserId, &res.TotalAmount, &res.Status, &res.CourierId, &res.Location, &res.DeliverySchedule)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("order not found")
		}
		return nil, err
	}
	return &res, nil
}

func (p *OrderRepo) GetAllOrder(flt *pb.GetAllOrdersRequest) (*pb.GetAllOrdersResponse, error) {
	limit := int32(10)
	offset := int32(0)

	if flt.Limit > 0{
		limit = flt.Limit
	}
	if flt.Page > 0{
		offset = (flt.Page - 1) * limit
	}
	query := `
		SELECT id,user_id,total_amount,status,courier_id,location,delivery_schedule
		FROM orders
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`
	rows, err := p.db.Query(query, limit, offset)
	if err != nil {
		log.Printf("Failed to get all orders: %v", err)
		return nil, fmt.Errorf("failed to get all orders: %v", err)
	}
	defer rows.Close()

	var orders []*pb.Order
	for rows.Next() {
		var res pb.Order
		if err := rows.Scan(
			&res.Id,
			&res.UserId,
			&res.TotalAmount,
			&res.Status,
			&res.CourierId,
			&res.Location,
			&res.DeliverySchedule,
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

	return &pb.GetAllOrdersResponse{
		Orders: orders,
	}, nil
}

func (p *OrderRepo) GetByLocation(order *pb.GetByLocationRequest)(*pb.GetByLocationResponse,error){
	query := `
		SELECT id,user_id,total_amount
		FROM orders WHERE location = $1
		ORDER BY created_at DESC
	`
	rows, err := p.db.Query(query,order.Location)
	if err != nil {
		log.Printf("Failed to get all orders: %v", err)
		return nil, fmt.Errorf("failed to get all orders: %v", err)
	}
	defer rows.Close()

	var orders []*pb.Location
	for rows.Next() {
		var res pb.Location
		if err := rows.Scan(
			&res.Id,
			&res.UserId,
			&res.TotalAmount,
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

	return &pb.GetByLocationResponse{
		Infos: orders,
	}, nil
}

func (p *OrderRepo) UpdateStatus(status *pb.UpdateStatusRequest)(*pb.UpdateStatusResponse, error){
	query := `UPDATE orders SET status = $1 WHERE id = $2`
	_, err := p.db.Exec(query,status.Status,status.OrderId)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (p *OrderRepo) GetHistory(history *pb.GetHistoryRequest)(*pb.GetHistoryResponse, error){
	limit := int32(10)
	offset := int32(0)

	if history.Limit > 0{
		limit = history.Limit
	}
	if history.Page > 0{
		offset = (history.Page - 1) * limit
	}
	query := `
		SELECT id,user_id,total_amount,status
		FROM orders WHERE status = 'Delivered'
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`
	rows, err := p.db.Query(query, limit, offset)
	if err != nil {
		log.Printf("Failed to get all orders: %v", err)
		return nil, fmt.Errorf("failed to get all orders: %v", err)
	}
	defer rows.Close()

	var orders []*pb.History
	for rows.Next() {
		var res pb.History
		if err := rows.Scan(
			&res.OrderId,
			&res.UserId,
			&res.TotalAmount,
			&res.Status,
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

	return &pb.GetHistoryResponse{
		History: orders,
	}, nil
}

