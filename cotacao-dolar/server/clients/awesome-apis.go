package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Docs: https://docs.awesomeapi.com.br/api-de-moedas
const API_URL = "https://economia.awesomeapi.com.br/json"

type CurrencyData struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

type Response struct {
	USDBRL CurrencyData `json:"USDBRL"`
}

type ExchangeAPI struct {
	Code   string
	CodeIn string
	CurrencyData
}

func NewExchangeClient(code, codeIn string) *ExchangeAPI {
	return &ExchangeAPI{
		Code:   code,
		CodeIn: codeIn,
	}
}

func (ex *ExchangeAPI) FetchLastExchange() error {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()

	url := fmt.Sprintf("%s/last/%s-%s", API_URL, ex.Code, ex.CodeIn)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Fatalf("Failed to create request for fetch exchange: %v", err)
		return err
	}

	body, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Failed to fetch exchange: %v", err)
		return err
	}

	var exchange Response
	json.NewDecoder(body.Body).Decode(&exchange)
	ex.CurrencyData = exchange.USDBRL

	return nil
}
