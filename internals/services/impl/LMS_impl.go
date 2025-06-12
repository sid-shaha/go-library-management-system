package impl

import (
	"github.com/sid-shaha/go-library-management-system/internals/models"
	"github.com/sid-shaha/go-library-management-system/internals/repository"
)

type LMSImpl struct {
	libraryRepo repository.LibraryRepo
}

func NewLMSService(libraryRepo repository.LibraryRepo) *LMSImpl {
	return &LMSImpl{
		libraryRepo: libraryRepo,
	}
}
func (l *LMSImpl) AddLibrary(library models.Library) (models.Library, error) {
	return l.libraryRepo.AddLibrary(library)
}
func (l *LMSImpl) DeleteLibrary(libraryID int) error {
	return l.libraryRepo.DeleteLibrary(libraryID)
}
func (l *LMSImpl) UpdateLibrary(libraryID int, newLibraryData models.Library) (models.Library, error) {
	return l.libraryRepo.UpdateLibrary(libraryID, newLibraryData)
}
func (l *LMSImpl) GetLibraryDetails(libraryID int) (models.Library, error) {
	return l.libraryRepo.GetLibraryDetails(libraryID)
}
