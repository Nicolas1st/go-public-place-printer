package dirsize

import (
	"os"
	"path/filepath"
)

func GetDirectorySizeInGB(path string) int64 {
	var dirSize int64 = 0
	readSize := func(path string, file os.FileInfo, err error) error {
		if !file.IsDir() {
			dirSize += file.Size()
		}
		return nil
	}

	filepath.Walk(path, readSize)
	sizeGB := dirSize / 1024 / 1024 / 1024

	return sizeGB
}
