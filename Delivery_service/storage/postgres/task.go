package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	pb "delivery/genprotos"
)

type TaskRepo struct {
	db *sql.DB
}

func NewTaskRepo(db *sql.DB) *TaskRepo {
	return &TaskRepo{db: db}
}

func (p *TaskRepo) CreateTask(task *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
	query := `INSERT INTO tasks(title, description, status, assigned_to,due_date)
	          VALUES($1, $2, $3, $4, $5)
	          RETURNING id, title, description, status, assigned_to, due_date`
	var res pb.CreateTaskResponse
	err := p.db.QueryRow(query, task.Title, task.Description, task.Status, task.AssignedTo, task.DueDate).Scan(
		&res.Id,
		&res.Title,
		&res.Description,
		&res.Status,
		&res.AssignedTo,
		&res.DueDate,
	)
	if err != nil {
		log.Printf("Failed to create task: %v", err)
		return nil, fmt.Errorf("failed to create task: %v", err)
	}

	return &res, nil
}

func (p *TaskRepo) UpdateTask(task *pb.UpdateTaskRequest) (*pb.UpdateTaskResponse, error) {
	query := `UPDATE tasks SET `
	var condition []string
	var args []interface{}

	if task.Title != "" && task.Title != "string" {
		condition = append(condition, fmt.Sprintf("title = $%d", len(args)+1))
		args = append(args, task.Title)
	}

	if task.Description != "" && task.Description != "string" {
		condition = append(condition, fmt.Sprintf("description = $%d", len(args)+1))
		args = append(args, task.Description)
	}
	if task.Status != "" && task.Status != "string" {
		condition = append(condition, fmt.Sprintf("status = $%d", len(args)+1))
		args = append(args, task.Status)
	}
	if task.AssignedTo != "" && task.AssignedTo != "string" {
		condition = append(condition, fmt.Sprintf("assigned_to = $%d", len(args)+1))
		args = append(args, task.AssignedTo)
	}
	if task.DueDate != "" && task.DueDate != "string" {
		condition = append(condition, fmt.Sprintf("due_date = $%d", len(args)+1))
		args = append(args, task.DueDate)
	}
	if len(condition) == 0 {
		return nil, errors.New("nothing to update")
	}

	query += strings.Join(condition, ", ")
	query += fmt.Sprintf(" WHERE id = $%d and deleted_at = 0 RETURNING id,title,description,assigned_to,due_date", len(args)+1)
	args = append(args, task.Id)

	res := pb.UpdateTaskResponse{}
	row := p.db.QueryRow(query, args...)
	err := row.Scan(&res.Id, &res.Title, &res.Description, &res.Status, &res.AssignedTo, &res.DueDate)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (p *TaskRepo) DeleteTask(id *pb.DeleteTaskRequest) (*pb.DeleteTaskResponse, error) {
	query := `UPDATE tasks SET deleted_at = EXTRACT(EPOCH FROM NOW()) WHERE id = $1 and deleted_at = 0`
	res, err := p.db.Exec(query, id.Id)
	if err != nil {
		log.Println("Error while delete task", err)
		return nil, err
	}
	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("task with id %s not found", id.Id)
	}
	log.Println("Successfully deleted task")

	return &pb.DeleteTaskResponse{}, nil
}

func (p *TaskRepo) GetByIdTask(id *pb.GetByIdTaskRequest) (*pb.GetByIdTaskResponse, error) {
	query := `SELECT id,title,description,status,assigned_to, due_date FROM tasks WHERE id = $1 and deleted_at = 0`
	row := p.db.QueryRow(query, id.Id)
	res := pb.GetByIdTaskResponse{}
	err := row.Scan(&res.Id, &res.Title, &res.Description, &res.Status, &res.AssignedTo, &res.DueDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("task not found")
		}
		return nil, err
	}
	return &res, nil
}

func (p *TaskRepo) GetAllTask(flt *pb.GetAllTasksRequest) (*pb.GetAllTasksResponse, error) {
	limit := int32(10)
	offset := int32(0)

	if flt.Limit > 0 {
		limit = flt.Limit
	}
	if flt.Page > 0 {
		offset = (flt.Page - 1) * limit
	}

	query := `
		SELECT id,title,description,status,assigned_to, due_date
		FROM tasks
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`
	rows, err := p.db.Query(query, limit, offset)
	if err != nil {
		log.Printf("Failed to get all tasks: %v", err)
		return nil, fmt.Errorf("failed to get all tasks: %v", err)
	}
	defer rows.Close()

	var tasks []*pb.GetByIdTaskResponse
	for rows.Next() {
		var task pb.GetByIdTaskResponse
		if err := rows.Scan(
			&task.Id,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.AssignedTo,
			&task.DueDate,
		); err != nil {
			log.Printf("Failed to scan task row: %v", err)
			return nil, fmt.Errorf("failed to scan task row: %v", err)
		}
		tasks = append(tasks, &task)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Failed to iterate over tasks: %v", err)
		return nil, fmt.Errorf("failed to iterate over tasks: %v", err)
	}

	return &pb.GetAllTasksResponse{
		Tasks: tasks,
	}, nil
}
