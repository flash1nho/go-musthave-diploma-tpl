package controllers

import (
    "fmt"
    "net/http"
    "encoding/json"

    "github.com/flash1nho/go-musthave-diploma-tpl/app/models"
    "github.com/flash1nho/go-musthave-diploma-tpl/middlewares"

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
        http.Error(w, "неверный формат запроса", http.StatusBadRequest)
        return
    }

    userID, err := models.UserRegister(requestBody.Login, requestBody.Password, controller.Pool)

    if err != nil {
        http.Error(w, "внутренняя ошибка сервера", http.StatusInternalServerError)
        controller.Log.Error(fmt.Sprint(err))
        return
    }

    if userID == 0 {
        w.WriteHeader(http.StatusConflict)
        fmt.Fprintln(w, "логин уже занят")
        return
    }

    err = setSignedCookie(userID, w)

    if err != nil {
        http.Error(w, "внутренняя ошибка сервера", http.StatusInternalServerError)
        controller.Log.Error(fmt.Sprint(err))
        return
    }

    w.WriteHeader(http.StatusOK)
    fmt.Fprintln(w, "пользователь успешно зарегистрирован и аутентифицирован")
}

func (controller *UserController) Login(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    var requestBody RequestBody

    err := json.NewDecoder(r.Body).Decode(&requestBody)

    if requestBody.Login == "" || requestBody.Password == "" {
        http.Error(w, "неверный формат запроса", http.StatusBadRequest)
        return
    }

    userID, err := models.UserLogin(requestBody.Login, requestBody.Password, controller.Pool)

    if err != nil {
        http.Error(w, "внутренняя ошибка сервера", http.StatusInternalServerError)
        controller.Log.Error(fmt.Sprint(err))
        return
    }

    if userID == 0 {
        w.WriteHeader(http.StatusUnauthorized)
        fmt.Fprintln(w, "неверная пара логин/пароль")
        return
    }

    err = setSignedCookie(userID, w)

    if err != nil {
        http.Error(w, "внутренняя ошибка сервера", http.StatusInternalServerError)
        controller.Log.Error(fmt.Sprint(err))
        return
    }

    w.WriteHeader(http.StatusOK)
    fmt.Fprintln(w, "пользователь успешно аутентифицирован")
}

func setSignedCookie(userID int, w http.ResponseWriter) error {
    encodedValue, err := middlewares.SecureCookieManager.Encode(middlewares.CookieName, userID)

    if err != nil {
        return err
    }

    cookie := &http.Cookie{
        Name:     middlewares.CookieName,
        Value:    encodedValue,
        Path:     "/",
        HttpOnly: true,
        Secure:   false,
        SameSite: http.SameSiteLaxMode,
        MaxAge:   3600 * 24 * 7,
    }

    http.SetCookie(w, cookie)

    return nil
}
