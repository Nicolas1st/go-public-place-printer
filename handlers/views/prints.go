package views

import (
	"fmt"
	"net/http"
	"printer/handlers/views/pages"
	"printer/persistence/model"
	"strconv"
)

func (c *viewsController) buildPrintsView(p *pages.Pages) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var daysAgo int
		{
			temp := r.URL.Query()["day"]
			daysAgo, _ = strconv.Atoi(temp[0])
		}

		// prints := c.db.GetAllPrints()
		prints, err := c.db.GetPrintsForDayNDaysAgo(daysAgo)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		days := []int{}
		for i := 0; i < 30; i++ {
			days = append(days, i)
		}

		fmt.Println(p.Admin.Prints.Execute(w, pages.NewFlashMessages(), struct {
			Prints []model.Print
			Days   []int
		}{
			Prints: prints,
			Days:   days,
		}))
	}
}
