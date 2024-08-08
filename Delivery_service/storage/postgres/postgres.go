package postgres

import (
	"database/sql"
	"fmt"
	"log"

	"delivery/config"
	postgres "delivery/storage"

	_ "github.com/lib/pq"
)

type Storage struct {
	db         *sql.DB
	CartC      postgres.CartI
	OrderItemC postgres.OrderItemsI
	OrderC     postgres.OrderI
	ProductC   postgres.ProductI
	TaskC      postgres.TaskI
	NotificationC postgres.NotificationI
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{db: db}
}

func ConnectDb() (*Storage, error) {
	cfg := config.Load()
	dbCon := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase)

		fmt.Println(dbCon)

	// Open database connection
	db, err := sql.Open("postgres", dbCon)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	cartS := NewCartRepo(db)
	orderitemS := NewOrderItemsRepo(db)
	orderS := NewOrderRepo(db)
	productS := NewProductRepo(db)
	taskS := NewTaskRepo(db)
	notificationS := NewNotificationRepo(db)

	return &Storage{
		CartC:      cartS,
		OrderItemC: orderitemS,
		OrderC:     orderS,
		ProductC:   productS,
		TaskC:      taskS,
		NotificationC: notificationS,
	}, nil
}

func (s *Storage) Cart() postgres.CartI {
	if s.CartC == nil {
		s.CartC = NewCartRepo(s.db)
	}
	return s.CartC
}

func (s *Storage) Order() postgres.OrderI {
	if s.OrderC == nil {
		s.OrderC = NewOrderRepo(s.db)
	}
	return s.OrderC
}

func (s *Storage) OrderItems() postgres.OrderItemsI {
	if s.OrderItemC == nil {
		s.OrderItemC = NewOrderItemsRepo(s.db)
	}
	return s.OrderItemC
}

func (s *Storage) Product() postgres.ProductI {
	if s.ProductC == nil {
		s.ProductC = NewProductRepo(s.db)
	}
	return s.ProductC
}

func (s *Storage) Task() postgres.TaskI {
	if s.TaskC == nil {
		s.TaskC = NewTaskRepo(s.db)
	}
	return s.TaskC
}

func (s *Storage) Notification() postgres.NotificationI{
	if s.NotificationC == nil{
		s.NotificationC = NewNotificationRepo(s.db)
	}
	return s.NotificationC
}