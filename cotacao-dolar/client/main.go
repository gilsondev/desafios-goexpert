package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type ExchangeData struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"var_bid"`
	PctChange  string `json:"pct_change"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/cotacao", nil)
	if err != nil {
		log.Fatalf("Failed to create request for fetch exchange: %v", err)
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Failed to fetch exchange: %v", err)
		panic(err)
	}
	defer res.Body.Close()

	var exchange ExchangeData
	json.NewDecoder(res.Body).Decode(&exchange)

	SaveExchangeToFile(exchange)
}

func SaveExchangeToFile(exchange ExchangeData) {

	file, err := os.Create("cotacao-dolar.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
		panic(err)
	}

	content := fmt.Sprintf("Dólar: %s", exchange.Bid)
	file.WriteString(string(content))
}
