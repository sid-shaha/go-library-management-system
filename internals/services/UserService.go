package services

import "github.com/sid-shaha/go-library-management-system/internals/models"

type UserService interface {
	addUser(libraryID, user models.User) (models.User, error)
	deleteUser(libraryID, userID int) error
	updateUser(libraryID int, userID int, NewUserData models.User) (models.User, error)
	getUserDetails(libraryID int, userID int) (models.User, error)
}
