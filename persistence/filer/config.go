package fileuploader

type config struct {
	pathToStoreFiles  string
	maxFileSize       int64
	formFieldFileName string
}

func NewConfig(pathToStoreFiles string) *config {
	return &config{
		pathToStoreFiles: pathToStoreFiles,
	}
}
