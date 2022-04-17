package jobq

import (
	"errors"
	"fmt"
	"printer/persistence/model"
)

type jobStatus byte

const (
	cancelled jobStatus = iota
	toBeDone
)

type JobQueue struct {
	JobIDGenerator *JobIDGenerator
	jobs           chan model.Job
	jobsStatus     map[model.JobID]jobStatus // to avoid linear search time
}

func NewJobQueue() *JobQueue {
	return &JobQueue{
		JobIDGenerator: newJobIDGenerator(),
		jobs:           make(chan model.Job, 20),
		jobsStatus:     make(map[model.JobID]jobStatus),
	}
}

func (q *JobQueue) Enqueue(job model.Job) model.JobID {
	jobID := job.ID

	// set the status for the job
	q.jobsStatus[jobID] = toBeDone

	// push the job on to the queue
	q.jobs <- job

	return jobID
}

func (q *JobQueue) Dequeue() (model.Job, error) {
	// get the next job
	// if queue is empty return error
	// otherwise loop till the job has status not equal to cancelled
	for {
		select {
		case job := <-q.jobs:
			if q.jobsStatus[job.ID] == cancelled {
				continue
			}

			delete(q.jobsStatus, job.ID)
			return job, nil
		default:
			return model.Job{}, errors.New("job queue is empty")
		}
	}
}

func (q *JobQueue) CancelJob(jobID model.JobID) error {
	// check if the job is currently in the queue
	// checking to avoid memory leak
	if status, ok := q.jobsStatus[jobID]; ok {
		if status == cancelled {
			return fmt.Errorf("the job %v had been already been canceled", jobID)
		}
		q.jobsStatus[jobID] = cancelled
		return nil
	}

	return fmt.Errorf("job with id %v is not in the queue", jobID)
}
