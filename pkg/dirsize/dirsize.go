package dirsize

import (
	"os"
	"path/filepath"
)

func DirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})

	size = size / 1024 / 1024 / 1024
	return size, err
}

// func GetDirectorySizeInGB(path string) int64 {
// 	var dirSize int64 = 0
// 	readSize := func(path string, file os.FileInfo, err error) error {
// 		if !file.IsDir() {
// 			dirSize += file.Size()
// 		}
// 		return nil
// 	}

// 	filepath.Walk(path, readSize)
// 	sizeGB := dirSize / 1024 / 1024 / 1024

// 	return sizeGB
// }
