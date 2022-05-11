package db

import "printer/persistence/model"

func (wrapper *Database) SavePrint(username, submittedFileName, storedFileName string, numberOfPages int) error {
	user := model.User{}

	if result := wrapper.db.Where("Name = ?", username).First(&user); result.Error != nil {
		return result.Error
	}

	print := model.Print{
		SubmittedFileName: submittedFileName,
		StoredFileName:    storedFileName,
		NumberOfPages:     numberOfPages,
		User:              user,
	}

	if result := wrapper.db.Create(&print); result.Error != nil {
		return result.Error
	}

	return nil
}

func (wrapper *Database) GetAllPrintsByUID(UID uint) []model.Print {
	var prints []model.Print
	wrapper.db.Where(model.Print{UserID: UID}).Find(&prints)

	return prints
}
