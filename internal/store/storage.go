package store

import (
	"context"
	"database/sql"
)

type Storage struct {
	Post interface {
		Creat(context.Context) error
	}
	Users interface {
		Creat(context.Context) error
	}
}

func NewPostgresStorage(db *sql.DB) Storage {
	return Storage{}

}
