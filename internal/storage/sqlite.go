package storage

import (
	"context"
	"database/sql"
	"errors"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/sanzhh/todo/internal/lib/e"
	"github.com/sanzhh/todo/internal/models"
	"github.com/sanzhh/todo/internal/ports"
)

type sqlite struct {
	db *sql.DB
}

func NewSQLite(db *sql.DB) *sqlite {
	return &sqlite{
		db: db,
	}
}

func (s *sqlite) Insert(ctx context.Context, data ports.InsertDTO) (err error) {
	var (
		errorMessage = "Storage.sqlite.Insert"
	)

	defer func() { err = e.WrapS(errorMessage, err) }()
	if data.Date == nil {
		date := time.Now().Format(models.TodoDateFormat)
		data.Date = &date
	}

	query := sq.Insert("").Into("todos").Columns("name", "description", "date").Values(data.Name, data.Description, data.Date)

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}

	res, err := s.db.ExecContext(ctx, sql, args...)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("failed to insert todo")
	}

	return nil
}

func (s *sqlite) Update(ctx context.Context, data ports.UpdateDTO) (err error) {
	var (
		errorMessage = "Storage.sqlite.Update"
	)
	defer func() { err = e.WrapS(errorMessage, err) }()

	setMap := sq.Eq{}
	if data.Name != nil {
		setMap["name"] = data.Name
	}

	if data.Description != nil {
		setMap["description"] = data.Description
	}

	if data.Date != nil {
		setMap["date"] = data.Date
	}

	if data.Done != nil {
		setMap["done"] = data.Done
	}

	sql, args, err := sq.Update("todos").SetMap(setMap).Where(sq.Eq{"id": data.ID}).ToSql()
	if err != nil {
		return err
	}

	res, err := s.db.ExecContext(ctx, sql, args...)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return ports.ErrTodoNotFound
	}

	return nil
}

func (s *sqlite) Delete(ctx context.Context, id string) (err error) {
	var (
		errorMessage = "Storage.sqlite.Delete"
	)
	defer func() { err = e.WrapS(errorMessage, err) }()

	sql, args, err := sq.Delete("todos").Where(sq.Eq{"id": id}).ToSql()
	if err != nil {
		return err
	}

	res, err := s.db.ExecContext(ctx, sql, args...)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return ports.ErrTodoNotFound
	}

	return nil
}

func (s *sqlite) Get(ctx context.Context, offset uint64, limit uint64) (_ []models.Todo, err error) {
	var (
		errorMessage = "Storage.sqlite.Get"
	)
	defer func() { err = e.WrapS(errorMessage, err) }()

	sql, args, err := sq.Select("*").From("todos").Limit(limit).Offset(offset).ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := s.db.QueryContext(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []models.Todo

	for rows.Next() {
		var todo models.Todo

		err := rows.Scan(
			&todo.ID,
			&todo.Name,
			&todo.Description,
			&todo.Date,
			&todo.Done,
		)

		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	return todos, nil
}

func (s *sqlite) GetByDate(ctx context.Context, date string) (_ []models.Todo, err error) {
	var (
		errorMessage = "Storage.sqlite.GetByDate"
	)

	defer func() { err = e.WrapS(errorMessage, err) }()

	sql, args, err := sq.Select("*").From("todos").Where(sq.Eq{"date": date}).ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := s.db.QueryContext(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []models.Todo

	for rows.Next() {
		var todo models.Todo

		err := rows.Scan(
			&todo.ID,
			&todo.Name,
			&todo.Description,
			&todo.Date,
			&todo.Done,
		)

		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	return todos, nil
}

func (s *sqlite) GetByID(ctx context.Context, id string) (_ *models.Todo, err error) {
	var (
		errorMessage = "Storage.sqlite.GetByID"
	)

	defer func() { err = e.WrapS(errorMessage, err) }()

	sql, args, err := sq.Select("*").From("todos").Where(sq.Eq{"id": id}).ToSql()
	if err != nil {
		return nil, err
	}

	var todo models.Todo

	if err := s.db.QueryRowContext(ctx, sql, args...).Scan(
		&todo.ID,
		&todo.Name,
		&todo.Description,
		&todo.Date,
		&todo.Done,
	); err != nil {
		return nil, err
	}

	return &todo, err
}
