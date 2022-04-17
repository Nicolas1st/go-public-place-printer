package jobs

import "printer/persistence/model"

type jobField struct {
	ID model.JobID `json:"jobID"`
}

type response struct {
	Job jobField
}

func newResponse(jobID model.JobID) response {
	return response{
		Job: jobField{ID: jobID},
	}
}
