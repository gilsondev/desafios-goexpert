package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gilsondev/desafios-goexpert/cotacao-dolar/server/clients"
	"github.com/gilsondev/desafios-goexpert/cotacao-dolar/server/models"
)

func GetExchange(w http.ResponseWriter, r *http.Request) {
	client := clients.NewExchangeClient("USD", "BRL")
	err := client.FetchLastExchange()
	if err != nil {
		http.Error(w, "Failed to fetch last exchange", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	exchange := models.ExchangeCurrency{
		Code:       client.Code,
		Codein:     client.CodeIn,
		Name:       client.Name,
		High:       client.High,
		Low:        client.Low,
		VarBid:     client.VarBid,
		PctChange:  client.PctChange,
		Bid:        client.Bid,
		Ask:        client.Ask,
		Timestamp:  client.Timestamp,
		CreateDate: client.CreateDate,
	}
	err = exchange.Save()
	if err != nil {
		http.Error(w, "Failed to fetch last exchange", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(exchange)
}
