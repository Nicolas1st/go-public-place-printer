package stats

import (
	"net/http"
	"printer/handlers"

	"github.com/gorilla/mux"
)

type database interface {
	GetPagesPrintedOverLastMonth() []int
}

type statsController struct {
	db database
}

func NewApi(db database) *mux.Router {
	c := &statsController{db: db}
	r := mux.NewRouter()
	r.HandleFunc(handlers.DefaultEndpoints.StatsApi, c.getUsage).Methods(http.MethodGet)

	return r
}
