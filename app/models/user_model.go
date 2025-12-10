package models

import (
	  "context"
	  "errors"

		"github.com/jackc/pgx/v5/pgxpool"
    "github.com/jackc/pgx/v5/pgconn"
    "github.com/jackc/pgerrcode"
		"golang.org/x/crypto/bcrypt"
)

type User struct {
    ID       int
    Login    string
    Password string
}

func UserRegisterWith(login string, password string, pool *pgxpool.Pool) (bool, error) {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

		if err != nil {
				return false, err
		}

    _, err = pool.Exec(context.TODO(), `INSERT INTO users (login, password) VALUES ($1, $2)`, login, hashedPassword)

		if err != nil {
        var pgErr *pgconn.PgError

        if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
        		return true, nil
        }

				return false, err
		}

		return false, nil
}
