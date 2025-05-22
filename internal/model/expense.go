package model

import "time"

type Expense struct {
	ID     int       `json:"id"`
	Title  string    `json:"title"`
	Amount float64   `json:"amount"`
	Date   time.Time `json:"date"`
}
