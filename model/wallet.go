package model

type BalanceCurrency struct {
	Balance Currency `json:"balance"`
}

type Currency struct {
	USD float64 `db:"usd" json:"USD"`
	RUB float64 `db:"rub" json:"RUB"`
	EUR float64 `db:"eur" json:"EUR"`
}
