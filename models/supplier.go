package models

import "time"

type Supplier struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Email     string    `gorm:"size:150;unique" json:"email"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Supplier) SupplierTable() string {
	return "suppliers" // The table name will now be 'suppliers'
}
