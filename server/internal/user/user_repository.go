package user

import (
	"context"
	"database/sql"
	"log"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type repository struct {
	db DBTX
}

func NewRepository(db DBTX) Repository {
	return &repository{db: db}
}

func (r *repository) CreateUser(ctx context.Context, user *User) (*User, error) {
	var lastInserted int

	// checkUser := User{}
	// q := "SELECT id, username, email FROM users WHERE username=$1"
	// err := r.db.QueryRowContext(ctx, q, user.Username, user.Email).Scan(&checkUser.ID, &checkUser.Username, &checkUser.Email,)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(checkUser.ID)
	// fmt.Println(checkUser.ID)
	// if checkUser.ID < 0 {
	// 	return &checkUser, nil
	// }

	query := "INSERT INTO users (username, email, password) VALUES ($1, $2, $3) returning id"
	err := r.db.QueryRowContext(ctx, query, user.Username, user.Email, user.Password).Scan(&lastInserted)
	if err != nil {
		log.Fatal(err)
	}
	user.ID = int64(lastInserted)
	return user, nil
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	u := User{}

	query := "SELECT id, username, email, password FROM users WHERE email = $1"
	err := r.db.QueryRowContext(ctx, query, email).Scan(&u.ID, &u.Username, &u.Email, &u.Password)
	if err != nil {
		return &User{}, err
	}

	return &u, nil
}