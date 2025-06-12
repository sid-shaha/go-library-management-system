package models

type BorrowingDetails struct {
	UserID     int
	BookCopyID int
	BorrowedAt int
	DueAt      int
	ReturnedAt int
	Fine       float64
}
