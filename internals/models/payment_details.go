package models

type PaymentDetails struct {
	ID            int
	TransactionID int
	UserID        int
	AmountPaid    float64
	PaymentFor    string // make it enum "dues" "membership"
}
