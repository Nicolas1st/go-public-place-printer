package filer

import (
	"io"
	"os"
	"path"
)

type filer struct {
	pathToStoreFiles string
	filePather       *FilePather
}

func NewFiler(pathToStoreFiles string, maxFileSize int64) *filer {
	return &filer{
		pathToStoreFiles: pathToStoreFiles,
		filePather:       newFilePather(pathToStoreFiles),
	}
}

func (f *filer) StoreFile(uploadedFile io.Reader, username, submittedFilename string) (pathToStoreFile string, err error) {
	pathToStoreFile = f.filePather.newFilePath(username)
	err = os.MkdirAll(pathToStoreFile, os.ModePerm)
	if err != nil {
		return "", err
	}

	filepath := path.Join(pathToStoreFile, submittedFilename)
	file, err := os.Create(filepath)
	if err != nil {
		return "", err
	}

	defer file.Close()

	_, err = io.Copy(file, uploadedFile)
	if err != nil {
		return "", err
	}

	return filepath, nil
}

func (f *filer) RemoveFile(filePath string) error {
	return os.Remove(filePath)
}
