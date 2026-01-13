package models

import (
	"time"
)

type Purchasing struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	SupplierID int       `json:"supplier_id"`
	UserID     int       `json:"user_id"`
	GrandTotal int       `json:"grand_total"`
	Date       time.Time `gorm:"column:date" json:"date"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (Purchasing) PurchasingTable() string {
	return "purchasings"
}
