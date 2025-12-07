package controllers

import (
    "net/http"
    "strconv"
    "github.com/flash1nho/go-musthave-diploma-tpl/app/models"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idStr)

    if err != nil {
        http.Error(w, "Некорректный ID пользователя", http.StatusBadRequest)
        return
    }

    user, err := models.FindUserByID(id)

    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusOK)
}
