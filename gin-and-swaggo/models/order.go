package models

import (
	"time"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type Order struct {
	ID      int `gorm:"primaryKey" json:"id"`
	CustomerName string `gorm:"not null" json:"customer_name"`
	OrderAt      time.Time `json:"order_at"`
	ItemsID       pq.Int64Array `gorm:"type:int[]" json:"items_id"` 
}