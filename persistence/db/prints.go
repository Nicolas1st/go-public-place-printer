package db

import "printer/persistence/model"

func (wrapper *Database) SavePrint(UID uint, submittedFileName, storedFileName string, numberOfPages int) error {
	user := model.User{}
	if result := wrapper.db.First(user, UID); result.Error != nil {
		return result.Error
	}

	print := model.Print{
		SubmittedFileName: submittedFileName,
		StoredFileName:    storedFileName,
		NumberOfPages:     numberOfPages,
	}

	if result := wrapper.db.Create(&print); result.Error != nil {
		return result.Error
	}

	return nil
}
