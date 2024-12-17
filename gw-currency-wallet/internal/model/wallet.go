package model

type Currency struct {
	USD float64 `db:"usd" json:"USD"`
	RUB float64 `db:"rub" json:"RUB"`
	EUR float64 `db:"eur" json:"EUR"`
}

type Operation struct {
	UserId   int
	Amount   float64 `json:"amount" binding:"required"`
	Currency string  `json:"currency" binding:"required"`
}
