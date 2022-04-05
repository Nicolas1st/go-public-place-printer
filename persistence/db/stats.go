package db

import "time"

func (wrapper *Database) GetNumberOfPagesPrinterByUser(userID uint) (int, error) {
	user := User{}
	if result := wrapper.db.First(user, userID); result.Error != nil {
		return 0, result.Error
	}

	prints := []Print{}
	wrapper.db.Model(&user).Association("Languages").Find(&prints)

	var pagesPrinted int
	monthAgo := time.Now().Add(-30 * (time.Hour * 24))
	for _, print := range prints {
		if print.CreatedAt.After(monthAgo) {
			pagesPrinted += print.NumberOfPages
		}
	}

	return pagesPrinted, nil
}
