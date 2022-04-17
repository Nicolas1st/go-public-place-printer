package filer

import (
	"path"
	"strconv"
)

type FileNamer struct {
	pathToStoreFiles string
	nextFileID       int
}

func newFileNamer(pathToStoreFiles string) *FileNamer {
	return &FileNamer{
		pathToStoreFiles: pathToStoreFiles,
		nextFileID:       0,
	}
}

func (f FileNamer) newFilePath(username, submittedFileName string) string {
	filePath := path.Join(f.pathToStoreFiles, username, strconv.Itoa(f.nextFileID), submittedFileName)
	f.nextFileID++

	return filePath
}
