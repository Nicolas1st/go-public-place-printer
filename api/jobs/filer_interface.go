package jobs

import "io"

type filerInterface interface {
	StoreFile(uploadedFile io.Reader, username, submittedFilename string) (filepath string, err error)
	RemoveFile(filePath string) error
}
