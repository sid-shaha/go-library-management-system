package models

type MembershipDetails struct {
	ID                    int
	UserID                int
	PaymentID             int
	Status                string // TODO: make it enum
	StartDate             int
	EndDate               int
	DepositPaid           float64
	CurrentDues           float64
	NumberOfBooksBorrowed int
	NumberOfBooksReturned int
}
