package models

type Book struct {
	ID              int
	Author          string
	Title           string
	PublicationDate int
	Isbn            string
	TotalCopies     int
	AvailableCopies int
}
