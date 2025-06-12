package services

type DuesService interface {
	GetPendingDues(userID string) (amount float64, err error)
	PayDues(userID string, amount float64) (paymentID int, err error)
}
