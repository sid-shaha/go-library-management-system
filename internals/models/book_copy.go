package models

type BookCopy struct {
	BookID int
	CopyID int
	// TODO: copyStatus BookCopyStatus
	Locator Locator
}
