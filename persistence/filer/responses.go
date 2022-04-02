package fileuploader

type resultMessage string

const (
	fileUploadedMessage     resultMessage = "File has been uploaded"
	fileTooBigMessage       resultMessage = "The file sent is too big, send a smaller one"
	wrongFileFormatMessage  resultMessage = "The application works only with pdf files"
	couldNotReadFileMessage resultMessage = "Could not read the file"
)

type response struct {
	ResultMessage resultMessage `json:"ResultMessage"`
	FileName      string        `json:"FileName"`
	FileFormat    string        `json:"FileFormat"`
	MaxFileSize   int64         `json:"FileSize"`
}

func newFileUploadedResponse(filename string) response {
	return response{
		ResultMessage: fileUploadedMessage,
		FileName:      filename,
	}
}

func newWrongFileFormatReponse(fileformat string) response {
	return response{
		ResultMessage: fileTooBigMessage,
		FileFormat:    fileformat,
	}
}

func newFileTooBigReponse(filesize int64) response {
	return response{
		ResultMessage: fileTooBigMessage,
		MaxFileSize:   filesize,
	}
}

func newCouldNotReadFileResponse() response {
	return response{
		ResultMessage: couldNotReadFileMessage,
	}
}
