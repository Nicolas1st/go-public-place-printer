package db

import (
	"printer/persistence/model"
	"time"
)

func (wrapper *Database) GetNumberOfPagesPrintedByUser(userID uint) (int, error) {
	user := model.User{}
	if result := wrapper.db.First(user, userID); result.Error != nil {
		return 0, result.Error
	}

	prints := []model.Print{}
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
