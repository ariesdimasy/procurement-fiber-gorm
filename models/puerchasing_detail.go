package models

import "time"

type PurchasingDetail struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	PurchasingID int       `json:"purchasing_id"`
	ItemID       int       `json:"item_id"`
	Qty          int       `json:"qty"`
	Subtotal     int       `json:"subtotal"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (PurchasingDetail) PurchasingDetailTable() string {
	return "purchasing_details"
}
