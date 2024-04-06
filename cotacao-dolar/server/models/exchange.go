package models

import (
	"context"
	"log"
	"time"

	"github.com/gilsondev/desafios-goexpert/cotacao-dolar/server/database"
)

type ExchangeCurrency struct {
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

func (ex *ExchangeCurrency) Save() error {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel()

	db, err := database.OpenConnection()
	if err != nil {
		log.Fatalf("Failed to open connection with database: %v", err)
		panic(err)
	}
	defer db.Close()

	stmt, err := db.PrepareContext(ctx,
		`INSERT INTO exchange (
			code,
			codein,
			name,
			high,
			low,
			var_bid,
			pct_change,
			bid,
			ask,
			timestamp,
			create_date
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
	)
	if err != nil {
		log.Fatalf("Failed to insert exchange: %v", err)
		return err
	}

	_, err = stmt.ExecContext(ctx,
		ex.Code,
		ex.Codein,
		ex.Name,
		ex.High,
		ex.Low,
		ex.VarBid,
		ex.PctChange,
		ex.Bid,
		ex.Ask,
		ex.Timestamp,
		ex.CreateDate,
	)
	if err != nil {
		log.Fatalf("Failed to insert exchange: %v", err)
		return err
	}

	return nil
}
