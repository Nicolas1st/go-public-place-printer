package db

import "printer/persistence/model"

func (wrapper *Database) SetPagesPerMonth(userID uint, pagesPerMonth int) error {
	user := model.User{}

	if result := wrapper.db.First(&user, userID); result.Error != nil {
		return result.Error
	}

	user.PagesPerMonth = pagesPerMonth
	if result := wrapper.db.Save(&user); result.Error != nil {
		return result.Error
	}

	return nil
}

func (wrapper *Database) AllowUsingPrinter(userID uint) error {
	user := model.User{}

	if result := wrapper.db.First(&user, userID); result.Error != nil {
		return result.Error
	}

	user.CanUsePrinter = true
	if result := wrapper.db.Save(&user); result.Error != nil {
		return result.Error
	}

	return nil
}

func (wrapper *Database) ForbidUsingPrinter(userID uint) error {
	user := model.User{}

	if result := wrapper.db.First(&user, userID); result.Error != nil {
		return result.Error
	}

	user.CanUsePrinter = false
	if result := wrapper.db.Save(&user); result.Error != nil {
		return result.Error
	}

	return nil
}
