package db

import (
	"context"

	"github.com/Phati/demo-load-balancer/domain"
	"github.com/jmoiron/sqlx"
)

type Storer interface {
	InsertUser(context context.Context, user *domain.User, db *sqlx.DB) (*domain.User, error)
	GetUser(context context.Context, id int, db *sqlx.DB) (*domain.User, error)
}

type PgStore struct {
}

func (p *PgStore) InsertUser(context context.Context, user *domain.User, db *sqlx.DB) (*domain.User, error) {
	sqlQuery := "Insert into users(id,email,pwd) values(nextval('users_sequence'),$1,$2) returning id,email"
	newUser := domain.User{}
	err := db.QueryRow(sqlQuery, user.Email, user.Password).Scan(&newUser.ID, &newUser.Email)
	if err != nil {
		return nil, err
	}

	return &newUser, nil
}

func (p *PgStore) GetUser(context context.Context, id int, db *sqlx.DB) (*domain.User, error) {
	sqlQuery := "Select * from users where id=$1"
	newUser := domain.User{}
	err := db.Get(&newUser, sqlQuery, id)
	if err != nil {
		return nil, err
	}

	return &newUser, nil
}
