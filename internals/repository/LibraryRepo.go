package repository

import "github.com/sid-shaha/go-library-management-system/internals/models"

type LibraryRepo interface {
	AddLibrary(library models.Library) (models.Library, error)
	DeleteLibrary(libraryID int) error
	UpdateLibrary(libraryID int, newLibraryData models.Library) (models.Library, error)
	GetLibraryDetails(libraryID int) (models.Library, error)
}
