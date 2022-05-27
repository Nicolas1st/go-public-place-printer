package views

import (
	"fmt"
	"net/http"
	"printer/handlers/views/pages"
	"printer/persistence/model"
)

func (c *viewsController) buildUserPrintsView(p *pages.Pages) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.URL.Query()["username"][0]
		prints := c.db.GetAllPrintsByUsername(username)

		fmt.Println(p.Admin.UserPrints.Execute(w, pages.NewFlashMessages(), struct {
			Prints   []model.Print
			Username string
		}{
			Prints:   prints,
			Username: username,
		}))
	}
}
