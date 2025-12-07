package main

import (
    "net/http"
    "github.com/flash1nho/go-musthave-diploma-tpl/app/controllers"
    "fmt"
)

func main() {
    http.HandleFunc("/user", controllers.GetUserHandler)

    fmt.Println("Сервер запущен на http://localhost:8080")

    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Ошибка запуска сервера:", err)
    }
}
