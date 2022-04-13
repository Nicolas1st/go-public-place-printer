package fileuploader

import (
	"io"
	"os"
)

type Filer struct {
	pathToStoreFiles string
	maxFileSize      int64
	storedFilePaths  map[string]string
	fileNamer        *FileNamer
}

func NewFiler(pathToStoreFiles string, maxFileSize int64) *Filer {
	return &Filer{
		pathToStoreFiles: pathToStoreFiles,
		maxFileSize:      maxFileSize,
		storedFilePaths:  map[string]string{},
		fileNamer:        newFileNamer(pathToStoreFiles),
	}
}

func (f *Filer) StoreFile(uploadedFile io.Reader, username string) (filename string, err error) {
	filename = f.fileNamer.newFilePath(username)

	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		return "", err
	}

	_, err = io.Copy(file, uploadedFile)
	if err != nil {
		return "", err
	}

	return filename, nil
}

func (f *Filer) RemoveFile(filePath string) error {
	if _, ok := f.storedFilePaths[filePath]; ok {
		return nil
	}

	return os.Remove(filePath)
}
