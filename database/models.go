// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"time"
)

type Stock struct {
	ID        int64
	Name      string
	Price     float32
	Company   string
	Createdat time.Time
	Updatedat time.Time
}
