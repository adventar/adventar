package db

import (
	"database/sql"

	"github.com/adventar/adventar/backend/pkg/gen/sqlc/adventar_db"
)

type Client interface {
	Queries() *adventar_db.Queries
	Close() error
}

type clientImpl struct {
	queries *adventar_db.Queries
	rawDb   *sql.DB
}

func New(dsn string) (Client, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	return &clientImpl{
		rawDb:   db,
		queries: adventar_db.New(db),
	}, nil
}

func (x *clientImpl) Queries() *adventar_db.Queries {
	return x.queries
}

func (x *clientImpl) Close() error {
	return x.rawDb.Close()
}
