package db

import (
	"fmt"
	"printer/persistence/model"
	"time"
)

func (wrapper *Database) SavePrint(user model.User, filename, filepath string, numberOfPages int) error {
	print := model.Print{
		NumberOfPages: numberOfPages,
		Filename:      filename,
		Filepath:      filepath,
		Username:      user.Name,
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

func (wrapper *Database) GetAllPrintsByUsername(username string) []model.Print {
	var prints []model.Print
	user, err := wrapper.GetUserByName(username)
	if err != nil {
		return prints
	}

	wrapper.db.Where(model.Print{UserID: user.ID}).Find(&prints)

	return prints
}

func (wrapper *Database) GetPrintsForDayNDaysAgo(daysAgo int) ([]model.Print, error) {
	var prints []model.Print

	startTime := time.Now().Add(-time.Duration(daysAgo+1) * (24 * time.Hour))
	endTime := time.Now().Add(-time.Duration(daysAgo) * (24 * time.Hour))
	result := wrapper.db.Where("created_at > ? AND created_at < ?", startTime, endTime).Find(&prints)

	return prints, result.Error
}

func (wrapper *Database) GetNumberOfPagesPrintedByUserDuringTheLastMonth(uid uint) int {
	var prints []model.Print
	const daysInMonth = 30

	startTime := time.Now().Add(-time.Duration(daysInMonth) * (24 * time.Hour))
	endTime := time.Now()
	wrapper.db.Where("created_at > ? AND created_at < ?", startTime, endTime).Where(model.Print{UserID: uid}).Find(&prints)

	total := 0
	for _, print := range prints {
		total += print.NumberOfPages
	}

	return total
}

func (wrapper *Database) GetPagesPrintedOverLastMonth() []int {
	currentTime := time.Now()
	days := 30

	var prints []model.Print
	wrapper.db.Where("created_at > ?", currentTime.Add(-time.Duration(days)*(24*time.Hour))).Find(&prints)
	fmt.Println(len(prints))
	for _, print := range prints {
		fmt.Println(print)
	}

	pagesPerDay := make([]int, days)
	for _, print := range prints {
		day := int(currentTime.Sub(print.CreatedAt).Hours()) % 24
		pagesPerDay[day] += print.NumberOfPages
	}

	return pagesPerDay
}
