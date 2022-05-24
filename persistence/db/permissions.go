package db

import "printer/persistence/model"

func (wrapper *Database) SetPagesPerMonth(userID uint, pagesPerMonth uint) (uint, error) {
	user := model.User{}

	if result := wrapper.db.First(&user, userID); result.Error != nil {
		return user.PagesPerMonth, result.Error
	}

	user.PagesPerMonth = pagesPerMonth
	if result := wrapper.db.Save(&user); result.Error != nil {
		return user.PagesPerMonth, result.Error
	}

	return user.PagesPerMonth, nil
}

func (wrapper *Database) AllowUsingPrinter(userID uint) (bool, error) {
	user := model.User{}

	if result := wrapper.db.First(&user, userID); result.Error != nil {
		return user.CanUsePrinter, result.Error
	}

	user.CanUsePrinter = true
	if result := wrapper.db.Save(&user); result.Error != nil {
		return user.CanUsePrinter, result.Error
	}

	return user.CanUsePrinter, nil
}

func (wrapper *Database) ForbidUsingPrinter(userID uint) (bool, error) {
	user := model.User{}

	if result := wrapper.db.First(&user, userID); result.Error != nil {
		return user.CanUsePrinter, result.Error
	}

	user.CanUsePrinter = false
	if result := wrapper.db.Save(&user); result.Error != nil {
		return user.CanUsePrinter, result.Error
	}

	return user.CanUsePrinter, nil
}
