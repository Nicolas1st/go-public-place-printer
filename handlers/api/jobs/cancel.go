package jobs

import (
	"net/http"
	"printer/persistence/model"
	"strconv"

	"github.com/gorilla/mux"
)

func (c *jobsController) CancelJob(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	c.jobq.CancelJob(model.JobID(id))
}
