package db

import "printer/persistence/model"

func (wrapper *Database) SavePrint(user model.User, filename string, numberOfPages int) error {
	print := model.Print{
		NumberOfPages: numberOfPages,
		Filename:      filename,
		User:          user,
	}

	result := wrapper.db.Create(&print)

	return result.Error
}

func (wrapper *Database) GetAllPrints() []model.Print {
	var prints []model.Print
	wrapper.db.Find(&prints)

	return prints
}

func (wrapper *Database) GetAllPrintsByUID(UID uint) []model.Print {
	var prints []model.Print
	wrapper.db.Where(model.Print{UserID: UID}).Find(&prints)

	return prints
}
