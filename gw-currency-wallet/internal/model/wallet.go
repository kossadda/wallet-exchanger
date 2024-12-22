package model

// Currency represents the current exchange rates for different currencies.
type Currency struct {
	// USD is the exchange rate for USD.
	USD float64 `db:"usd" json:"USD"`

	// RUB is the exchange rate for RUB.
	RUB float64 `db:"rub" json:"RUB"`

	// EUR is the exchange rate for EUR.
	EUR float64 `db:"eur" json:"EUR"`
}

// Operation represents a financial transaction involving a user's wallet, such as a deposit or withdrawal.
type Operation struct {
	// UserId is the identifier of the user performing the operation.
	UserId int

	// Amount is the amount involved in the operation (e.g., the amount to be deposited or withdrawn).
	Amount float64 `json:"amount" binding:"required"`

	// Currency is the type of currency involved in the operation (e.g., USD, EUR).
	Currency string `json:"currency" binding:"required"`
}
