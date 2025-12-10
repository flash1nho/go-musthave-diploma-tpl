package handler

import (
	  "github.com/flash1nho/go-musthave-diploma-tpl/app/controllers"

		"github.com/jackc/pgx/v5/pgxpool"

		"go.uber.org/zap"
)

type Handler struct {
    UserController *controllers.UserController
    Log   *zap.Logger
}

func NewHandler(pool *pgxpool.Pool, log *zap.Logger) *Handler {
    return &Handler{
        UserController: &controllers.UserController{Pool: pool, Log: log},
        Log:   log,
    }
}
