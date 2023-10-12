package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Database struct {
	db	*sql.DB
}

func NewDatabase() (*Database, error) {
	dns := "postgres://root:password@localhost:5433/go-chat?sslmode=disable"
	db, err := sql.Open("postgres", dns)
	if err != nil {
		return nil, err
	}
	return &Database{db: db}, nil
}

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}