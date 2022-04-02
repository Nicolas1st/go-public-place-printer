package jobs

import "printer/interfaces"

type jobField struct {
	ID interfaces.JobID `json:"jobID"`
}

type Response struct {
	Job jobField
}

func NewResponse(jobID interfaces.JobID) Response {
	return Response{
		Job: jobField{ID: jobID},
	}
}
