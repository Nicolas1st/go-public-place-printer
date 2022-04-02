package fileuploader

import (
	"path"
	"strconv"
)

func newFileNamer(c config) func(username string) (filename string) {
	fileID := 0
	return func(username string) (filename string) {
		fileID++
		filename = path.Join(c.pathToStoreFiles, username, strconv.Itoa(fileID))
		return filename
	}
}
