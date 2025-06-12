package models

type User struct {
	ID            int
	Name          string
	Surname       string
	Password      string
	Email         string
	ContactNumber int
	UserRole      string //TODO: make it enum
}
