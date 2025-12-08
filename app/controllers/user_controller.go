package controllers

import (
    "net/http"

    "github.com/flash1nho/go-musthave-diploma-tpl/app/models"

    "github.com/jackc/pgx/v5/pgxpool"
)

type UserController struct {
    Pool *pgxpool.Pool
}

func (controller *UserController) Register(w http.ResponseWriter, r *http.Request) {
    login := r.URL.Query().Get("login")
    password := r.URL.Query().Get("password")
    _, err := models.UserRegisterWith(controller.Pool, login, password)

    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusOK)
}
