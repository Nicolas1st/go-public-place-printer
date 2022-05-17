package jobs

import (
	"encoding/json"
	"net/http"
)

func (c *jobsController) GetAllJobs(w http.ResponseWriter, r *http.Request) {
	jobs := c.jobq.GetAllJobs()
	json.NewEncoder(w).Encode(jobs)
}
