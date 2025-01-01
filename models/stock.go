package models

import(
	"time"
)

type Stock struct {
	ID        int64   `json:"id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Company   string  `json:"company"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}