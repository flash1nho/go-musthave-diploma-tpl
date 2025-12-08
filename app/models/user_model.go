package models

import (
		"fmt"

		"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
    ID       int
    Login    string
    Password string
}

func UserRegisterWith(pool *pgxpool.Pool, login string, password string) (*User, error) {
    if login == "login" {
        return &User{ID: 1, Login: "login", Password: "password"}, nil
    }

    return nil, fmt.Errorf("пользователь с login %s не найден", login)
}
