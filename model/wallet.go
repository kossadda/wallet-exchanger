package model

type BalanceCurrency struct {
	Balance Currency `json:"balance"`
}

type Currency struct {
	USD float64 `json:"USD"`
	RUB float64 `json:"RUB"`
	EUR float64 `json:"EUR"`
}
