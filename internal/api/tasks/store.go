package tasks

import (
	"context"
	"database/sql"
	"fmt"
)

type SQLStore struct {
	db *sql.DB
}

func NewSQLStore(db *sql.DB) Store {
	return &SQLStore{db: db}
}

var _ Store = (*SQLStore)(nil)

func (s *SQLStore) List(ctx context.Context) ([]TaskDTO, error) {
	rows, err := s.db.QueryContext(ctx, `
			SELECT id, description, done
			FROM tasks
			ORDER BY id`)
	if err != nil {
		return nil, fmt.Errorf("list tasks: %w", err)
	}
	defer rows.Close()

	var out []TaskDTO
	for rows.Next() {
		var t TaskDTO
		if err := rows.Scan(&t.Id, &t.Description, &t.Done); err != nil {

		}
		out = append(out, t)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate tasks: %w", err)
	}
	return out, nil
}

func (s *SQLStore) Create(ctx context.Context, description string) (TaskDTO, error) {
	var t TaskDTO
	err := s.db.QueryRowContext(ctx, `
		INSERT INTO tasks (description, done)
		VALUES ($1, false)
		RETURNING id, description, done`,
		description,
	).Scan(&t.Id, &t.Description, &t.Done)
	if err != nil {
		return TaskDTO{}, fmt.Errorf("create task, %w", err)
	}
	return t, nil
}
