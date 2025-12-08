package main

import (
    "fmt"

    "github.com/flash1nho/go-musthave-diploma-tpl/config"
    "github.com/flash1nho/go-musthave-diploma-tpl/db"
    "github.com/flash1nho/go-musthave-diploma-tpl/handler"
    "github.com/flash1nho/go-musthave-diploma-tpl/service"
)

func main() {
	  apiServer, accrualServer, log, databaseURI := config.Settings()
	  dbPool, err := db.NewDB(databaseURI)

    if err != nil {
        log.Error(fmt.Sprint(err))
    }

    h := handler.NewHandler(dbPool, log)
    servers := []config.Server{apiServer, accrualServer}
    service.NewService(h, servers).Run()
}
