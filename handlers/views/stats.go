package views

import (
	"fmt"
	"net/http"
	"printer/handlers/views/pages"
)

func (c *viewsController) buildStatsView(p *pages.Pages) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(p.Admin.Stats.Execute(w, pages.NewFlashMessages(), nil))
	}
}
