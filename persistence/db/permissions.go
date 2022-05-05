package db

import "printer/persistence/model"

func (wrapper *Database) SetUsersPagesPerMonth(id uint, pagesPerMonth int) error {
	user := model.User{}

	if result := wrapper.db.First(user, id); result.Error != nil {
		return result.Error
	}

	user.PagesPerMonth = pagesPerMonth
	if result := wrapper.db.Save(&user); result.Error != nil {
		return result.Error
	}

	return nil
}

func (wrapper *Database) AllowUserToUserPrinter(id uint) error {
	user := model.User{}

	if result := wrapper.db.First(user, id); result.Error != nil {
		return result.Error
	}

	user.CanUsePrinter = true
	if result := wrapper.db.Save(&user); result.Error != nil {
		return result.Error
	}

	return nil
}

func (wrapper *Database) ForbidUsingPrinter(id uint) error {
	user := model.User{}

	if result := wrapper.db.First(user, id); result.Error != nil {
		return result.Error
	}

	user.CanUsePrinter = true
	if result := wrapper.db.Save(&user); result.Error != nil {
		return result.Error
	}

	return nil
}
