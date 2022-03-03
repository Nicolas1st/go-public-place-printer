package jobs

import "printer/interfaces"

type CancelJobRequest struct {
	ID interfaces.JobID `json:"ID"`
}
