package impl

import "github.com/sid-shaha/go-library-management-system/internals/models"

// TODO: implement all the functions
type LibraryRepoImpl struct {
	//TODO: map id to libraryStruct
}

func (l *LibraryRepoImpl) AddLibrary(library models.Library) (models.Library, error) {
	return models.Library{}, nil
}
func (l *LibraryRepoImpl) DeleteLibrary(libraryID int) error {
	return nil
}
func (l *LibraryRepoImpl) UpdateLibrary(libraryID int, newLibraryData models.Library) (models.Library, error) {
	return models.Library{}, nil
}
func (l *LibraryRepoImpl) GetLibraryDetails(libraryID int) (models.Library, error) {
	return models.Library{}, nil
}
