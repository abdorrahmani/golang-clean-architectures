package models

type Product struct {
	ID          int     `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string  `json:"name" gorm:"size:255;not null"`
	Description string  `json:"description" gorm:"type:text"`
	Price       float64 `json:"price" gorm:"type:decimal(10,2);not null"`
	CreatedAt   string  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   string  `json:"updated_at" gorm:"autoUpdateTime"`
}
