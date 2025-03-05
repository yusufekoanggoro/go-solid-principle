package entity

type Payment struct {
	ID     string
	Amount float64
	Method string // Misalnya: "bank_transfer", "credit_card"
}
