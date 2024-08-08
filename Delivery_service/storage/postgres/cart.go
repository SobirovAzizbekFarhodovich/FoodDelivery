package postgres

import (
	"database/sql"
	pb "delivery/genprotos"
	"errors"
	"fmt"
	"log"
	"strings"
)

type CartRepo struct {
	db *sql.DB
}

func NewCartRepo(db *sql.DB) *CartRepo {
	return &CartRepo{db: db}
}

func (p *CartRepo) CreateCart(cart *pb.CreateCartRequest) (*pb.CreateCartResponse, error) {
	query := `INSERT INTO cart(user_id,product_id,quantity,options)
	          VALUES($1, $2, $3, $4)
	          RETURNING id, user_id,product_id,quantity,options`
	var res pb.CreateCartResponse
	err := p.db.QueryRow(query, cart.UserId, cart.ProductId, cart.Quantity, cart.Option).Scan(
		&res.Id,
		&res.UserId,
		&res.ProductId,
		&res.Quantity,
		&res.Option,
	)
	if err != nil {
		log.Printf("Failed to create order: %v", err)
		return nil, fmt.Errorf("failed to create order: %v", err)
	}

	return &res, nil
}

func (p *CartRepo) UpdateCart(order *pb.UpdateCartRequest) (*pb.UpdateCartResponse, error) {
	query := `UPDATE cart SET `
	var condition []string
	var args []interface{}

	if order.UserId != "" && order.UserId != "string" {
		condition = append(condition, fmt.Sprintf("user_id = $%d", len(args)+1))
		args = append(args, order.UserId)
	}

	if order.ProductId != "" && order.ProductId != "string" {
		condition = append(condition, fmt.Sprintf("product_id = $%d", len(args)+1))
		args = append(args, order.ProductId)
	}
	if order.Quantity != 0 {
		condition = append(condition, fmt.Sprintf("quantity = $%d", len(args)+1))
		args = append(args, order.Quantity)
	}
	if order.Option != "" && order.Option != "string" {
		condition = append(condition, fmt.Sprintf("options = $%d", len(args)+1))
		args = append(args, order.Option)
	}
	if len(condition) == 0 {
		return nil, errors.New("nothing to update")
	}

	query += strings.Join(condition, ", ")
	query += fmt.Sprintf(" WHERE id = $%d and deleted_at = 0 RETURNING id,user_id,product_id,quantity,options", len(args)+1)
	args = append(args, order.Id)

	res := pb.UpdateCartResponse{}
	row := p.db.QueryRow(query, args...)
	err := row.Scan(&res.Id, &res.UserId, &res.ProductId, &res.Quantity, &res.Option)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (p *CartRepo) DeleteCart(id *pb.DeleteCartRequest) (*pb.DeleteCartResponse, error) {
	query := `UPDATE cart SET deleted_at = EXTRACT(EPOCH FROM NOW()) WHERE id = $1 and deleted_at = 0`
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

	return &pb.DeleteCartResponse{}, nil
}

func (p *CartRepo) GetByIdCart(id *pb.GetByIdCartRequest) (*pb.GetByIdCartResponse, error) {
	query := `SELECT id,user_id,product_id,quantity,options FROM cart WHERE id = $1 and deleted_at = 0`
	row := p.db.QueryRow(query, id.Id)
	res := pb.GetByIdCartResponse{}
	err := row.Scan(&res.Id, &res.UserId, &res.ProductId, &res.Quantity, &res.Option)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("order not found")
		}
		return nil, err
	}
	return &res, nil
}

func (p *CartRepo) GetAllCart(flt *pb.GetAllCartRequest) (*pb.GetAllCartResponse, error) {
	limit := int32(10)
	offset := int32(0)

	if flt.Limit > 0 {
		limit = flt.Limit
	}
	if flt.Page > 0 {
		offset = (flt.Page - 1) * limit
	}

	query := `
		SELECT id, user_id, product_id, quantity, options
		FROM cart
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`
	rows, err := p.db.Query(query, limit, offset)
	if err != nil {
		log.Printf("Failed to get all carts: %v", err)
		return nil, fmt.Errorf("failed to get all carts: %v", err)
	}
	defer rows.Close()

	var carts []*pb.GetByIdCartResponse
	for rows.Next() {
		var res pb.GetByIdCartResponse
		if err := rows.Scan(
			&res.Id,
			&res.UserId,
			&res.ProductId,
			&res.Quantity,
			&res.Option,
		); err != nil {
			log.Printf("Failed to scan cart row: %v", err)
			return nil, fmt.Errorf("failed to scan cart row: %v", err)
		}
		carts = append(carts, &res)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Failed to iterate over carts: %v", err)
		return nil, fmt.Errorf("failed to iterate over carts: %v", err)
	}

	return &pb.GetAllCartResponse{
		Carts: carts,
	}, nil
}

