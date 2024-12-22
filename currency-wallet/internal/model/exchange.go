// Package model defines the data structures used for parsing and handling JSON input and output,
// including user data, wallet operations, and other domain-specific entities.
package model

// Exchange represents the data for a currency exchange transaction.
type Exchange struct {
	// UserId is the identifier of the user initiating the exchange.
	UserId int

	// FromCurrency is the currency that the user is exchanging from (e.g., USD).
	FromCurrency string `json:"from_currency" binding:"required"`

	// ToCurrency is the currency that the user is exchanging to (e.g., EUR).
	ToCurrency string `json:"to_currency" binding:"required"`

	// Amount is the amount of money to be exchanged.
	Amount float64 `json:"amount" binding:"required"`
}
