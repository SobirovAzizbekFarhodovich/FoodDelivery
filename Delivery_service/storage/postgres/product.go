package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	pb "delivery/genprotos"
)

type ProductRepo struct {
	db *sql.DB
}

func NewProductRepo(db *sql.DB) *ProductRepo {
	return &ProductRepo{db: db}
}

func (p *ProductRepo) CreateProduct(product *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	query := `INSERT INTO products(name, description, price, image)
	          VALUES($1, $2, $3, $4)
	          RETURNING id, name, description, price, image`
	var createdProduct pb.CreateProductResponse
	err := p.db.QueryRow(query, product.Name, product.Description, product.Price, product.Image).Scan(
		&createdProduct.Id,
		&createdProduct.Name,
		&createdProduct.Description,
		&createdProduct.Price,
		&createdProduct.Image,
	)
	if err != nil {
		log.Printf("Failed to create product: %v", err)
		return nil, fmt.Errorf("failed to create product: %v", err)
	}

	return &createdProduct, nil
}

func (p *ProductRepo) UpdateProduct(product *pb.UpdateProductRequest) (*pb.UpdateProductResponse, error) {
	query := `UPDATE products SET `
	var condition []string
	var args []interface{}

	if product.Name != "" && product.Name != "string" {
		condition = append(condition, fmt.Sprintf("name = $%d", len(args)+1))
		args = append(args, product.Name)
	}

	if product.Description != "" && product.Description != "string" {
		condition = append(condition, fmt.Sprintf("description = $%d", len(args)+1))
		args = append(args, product.Description)
	}
	if product.Price != 0 {
		condition = append(condition, fmt.Sprintf("price = $%d", len(args)+1))
		args = append(args, product.Price)
	}
	if product.Image != "" && product.Image != "string" {
		condition = append(condition, fmt.Sprintf("image = $%d", len(args)+1))
		args = append(args, product.Image)
	}
	if len(condition) == 0 {
		return nil, errors.New("nothing to update")
	}

	query += strings.Join(condition, ", ")
	query += fmt.Sprintf(" WHERE id = $%d and deleted_at = 0 RETURNING id,name,description,price,image", len(args)+1)
	args = append(args, product.Id)

	res := pb.UpdateProductResponse{}
	row := p.db.QueryRow(query, args...)
	err := row.Scan(&res.Id, &res.Name, &res.Description, &res.Price, &res.Image)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (p *ProductRepo) DeleteProduct(id *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error) {
	query := `UPDATE products SET deleted_at = EXTRACT(EPOCH FROM NOW()) WHERE id = $1 and deleted_at = 0`
	res, err := p.db.Exec(query, id.Id)
	if err != nil {
		log.Println("Error while delete product", err)
		return nil, err
	}
	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("product with id %s not found", id.Id)
	}
	log.Println("Successfully deleted product")

	return &pb.DeleteProductResponse{}, nil
}

func (p *ProductRepo) GetByIdProduct(id *pb.GetByIdProductRequest) (*pb.GetByIdProductResponse, error) {
	query := `SELECT id,name,description,price,image FROM products WHERE id = $1 and deleted_at = 0`
	row := p.db.QueryRow(query, id.Id)
	res := pb.GetByIdProductResponse{}
	err := row.Scan(&res.Id, &res.Name, &res.Description, &res.Price, &res.Image)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("product not found")
		}
		return nil, err
	}
	return &res, nil
}

func (p *ProductRepo) GetAllProduct(flt *pb.GetAllProductsRequest) (*pb.GetAllProductsResponse, error) {
	limit := int32(10)
	offset := int32(0)

	if flt.Limit > 0 {
		limit = flt.Limit
	}
	if flt.Page > 0 {
		offset = (flt.Page - 1) * limit
	}

	query := `
		SELECT id, name, description, price, image
		FROM products
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`
	rows, err := p.db.Query(query, limit, offset)
	if err != nil {
		log.Printf("Failed to get all products: %v", err)
		return nil, fmt.Errorf("failed to get all products: %v", err)
	}
	defer rows.Close()

	var products []*pb.GetByIdProductResponse
	for rows.Next() {
		var product pb.GetByIdProductResponse
		if err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Image,
		); err != nil {
			log.Printf("Failed to scan product row: %v", err)
			return nil, fmt.Errorf("failed to scan product row: %v", err)
		}
		products = append(products, &product)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Failed to iterate over products: %v", err)
		return nil, fmt.Errorf("failed to iterate over products: %v", err)
	}

	return &pb.GetAllProductsResponse{
		Products: products,
	}, nil
}
