package models

import "fmt"

type User struct {
    ID   int
    Name string
    Email string
}

func FindUserByID(id int) (*User, error) {
    if id == 1 {
        return &User{ID: 1, Name: "Иван Петров", Email: "ivan@example.com"}, nil
    }

    return nil, fmt.Errorf("пользователь с ID %d не найден", id)
}
