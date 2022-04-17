package jobs

import (
	"printer/persistence/model"
)

type cancelJobRequest struct {
	ID model.JobID `json:"ID"`
}
