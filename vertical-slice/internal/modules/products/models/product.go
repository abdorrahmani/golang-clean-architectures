package models

import "time"

type Product struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string    `json:"name" gorm:"size:255;not null"`
	Description string    `json:"description" gorm:"type:text"`
	Price       float64   `json:"price" gorm:"type:decimal(10,2);not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
