package postgres

import (
	"database/sql"
	pb "delivery/genprotos"
	"fmt"
	"log"
)

type NotificationRepo struct {
	db *sql.DB
}

func NewNotificationRepo(db *sql.DB) *NotificationRepo {
	return &NotificationRepo{db: db}
}

func (n *NotificationRepo) CreateNotification(not *pb.CreateNotificationRequest) (*pb.CreateNotificationResponse, error) {
	query := `
	INSERT INTO notification (user_id,courier_id, order_id, message)VALUES($1,$2,$3,$4)
	`
	_, err := n.db.Exec(query, not.UserId, not.CourierId, not.OrderId, not.Message)
	if err != nil {
		return nil, err
	}
	return nil, nil

}

func (n *NotificationRepo) GetNotification(flt *pb.GetNotificationsRequest) (*pb.GetNotificationsResponse, error) {
	limit := int32(10)
	offset := int32(0)

	if flt.Limit > 0 {
		limit = flt.Limit
	}
	if flt.Page > 0 {
		offset = (flt.Page - 1) * limit
	}
	query := `
		SELECT id,user_id,courier_id,order_id,message
		FROM notification
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`
	rows, err := n.db.Query(query, limit, offset)
	if err != nil {
		log.Printf("Failed to get all orders: %v", err)
		return nil, fmt.Errorf("failed to get all orders: %v", err)
	}
	defer rows.Close()

	var orders []*pb.Notification
	for rows.Next() {
		var res pb.Notification
		if err := rows.Scan(
			&res.Id,
			&res.UserId,
			&res.CourierId,
			&res.OrderId,
			&res.Message,
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

	return &pb.GetNotificationsResponse{
		Notifications: orders,
	}, nil
}
