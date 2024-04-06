package main

import (
	"net/http"

	"github.com/gilsondev/desafios-goexpert/cotacao-dolar/server/database"
	"github.com/gilsondev/desafios-goexpert/cotacao-dolar/server/handlers"
)

func main() {
	mux := http.NewServeMux()
	InitConfig()

	mux.HandleFunc("GET /cotacao", handlers.GetExchange)
	http.ListenAndServe(":8080", mux)
}

func InitConfig() {
	err := database.CreateSchema()
	if err != nil {
		panic(err)
	}
}
