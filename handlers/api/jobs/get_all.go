package jobs

import (
	"encoding/json"
	"net/http"
	"printer/handlers"
	"printer/persistence/model"
)

func (c *jobsController) GetAllJobs(w http.ResponseWriter, r *http.Request) {
	session, ok := handlers.GetSession(c.sessioner, r)
	if !ok {
		http.Redirect(w, r, handlers.DefaultEndpoints.LoginPage, http.StatusSeeOther)
		return
	}

	jobs := c.jobq.GetAllJobs()

	usersJobs := []*model.Job{}
	for _, job := range jobs {
		if job.User.ID == session.User.ID {
			usersJobs = append(usersJobs, job)
		}
	}

	json.NewEncoder(w).Encode(usersJobs)
}
