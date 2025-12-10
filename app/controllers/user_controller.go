package controllers

import (
    "fmt"
    "net/http"
    "encoding/json"

    "github.com/flash1nho/go-musthave-diploma-tpl/app/models"

    "github.com/jackc/pgx/v5/pgxpool"
    "go.uber.org/zap"
)

type UserController struct {
    Pool *pgxpool.Pool
    Log  *zap.Logger
}

type RequestBody struct {
    Login    string `json:"login"`
    Password string `json:"password"`
}

func (controller *UserController) Register(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    var requestBody RequestBody

    err := json.NewDecoder(r.Body).Decode(&requestBody)

    if requestBody.Login == "" || requestBody.Password == "" {
        http.Error(w, "Неверный формат данных", http.StatusBadRequest)
        return
    }

    userExists, err := models.UserRegisterWith(requestBody.Login, requestBody.Password, controller.Pool)

    if err != nil {
        http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
        controller.Log.Error(fmt.Sprint(err))
        return
    }

    if userExists {
        w.WriteHeader(http.StatusConflict)
        fmt.Fprintf(w, "Пользователь %s существует\n", requestBody.Login)
        return
    }

    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Пользователь %s успешно зарегистрирован\n", requestBody.Login)
}

// func (controller *UserController) Login(w http.ResponseWriter, r *http.Request) {
//     w.Header().Set("Content-Type", "application/json")

//     // TODO

//     w.WriteHeader(http.StatusOK)
// }
