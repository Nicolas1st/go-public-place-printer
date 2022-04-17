package filer

import (
	"io"
	"os"
)

type filer struct {
	pathToStoreFiles string
	fileNamer        *FileNamer
}

func NewFiler(pathToStoreFiles string, maxFileSize int64) *filer {
	return &filer{
		pathToStoreFiles: pathToStoreFiles,
		fileNamer:        newFileNamer(pathToStoreFiles),
	}
}

func (f *filer) StoreFile(uploadedFile io.Reader, username, submittedFilename string) (filepath string, err error) {
	filepath = f.fileNamer.newFilePath(username, submittedFilename)

	file, err := os.Create(filepath)
	defer file.Close()
	if err != nil {
		return "", err
	}

	_, err = io.Copy(file, uploadedFile)
	if err != nil {
		return "", err
	}

	return filepath, nil
}

func (f *filer) RemoveFile(filePath string) error {
	return os.Remove(filePath)
}
