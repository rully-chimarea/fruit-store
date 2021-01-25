package model

type Fruits struct {
	Id        uint    `gorm:"primaryKey"`
	FruitName string  `gorm:"column:fruit-name" json:"fruit-name"`
	FruitType string  `gorm:"column:fruit-type" json:"fruit-type"`
	Price     float64 `json:"price" gorm:"type:numeric"`
	Quantity  uint    `json:"quantity" `
}
