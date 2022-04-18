package filer

import (
	"path"
	"strconv"
)

type FilePather struct {
	pathToStoreFiles string
	nextFileID       int
}

func newFilePather(pathToStoreFiles string) *FilePather {
	return &FilePather{
		pathToStoreFiles: pathToStoreFiles,
		nextFileID:       0,
	}
}

func (f *FilePather) newFilePath(username string) string {
	filepath := path.Join(f.pathToStoreFiles, username, strconv.Itoa(f.nextFileID))
	f.nextFileID++

	return filepath
}
