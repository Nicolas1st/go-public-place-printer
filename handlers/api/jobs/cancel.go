package jobs

import (
	"net/http"
	"printer/persistence/model"

	"github.com/gorilla/mux"
)

func (c *jobsController) CancelJob(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	c.jobq.CancelJob(model.JobID(id))
}
