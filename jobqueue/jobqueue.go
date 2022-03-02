package jobqueue

import (
	"errors"
	"fmt"
	"printer/interfaces"
)

type jobStatus byte

const (
	cancelled jobStatus = iota
	toBeDone
)

type JobQueue struct {
	UniqueIDGenerator func() interfaces.JobID
	jobs              chan interfaces.Job
	jobsStatus        map[interfaces.JobID]jobStatus // to avoid linear search time
}

func NewJobQueue() *JobQueue {
	return &JobQueue{
		UniqueIDGenerator: NewJobIDGenerator(),
		jobs:              make(chan interfaces.Job),
		jobsStatus:        make(map[interfaces.JobID]jobStatus),
	}
}

func (q *JobQueue) Enqueue(job interfaces.Job) interfaces.JobID {
	// set unique id to incoming job
	jobID := q.UniqueIDGenerator()
	job.SetID(jobID)

	// set the status for the job
	q.jobsStatus[jobID] = toBeDone

	// push the job on to the queue
	q.jobs <- job

	return jobID
}

func (q *JobQueue) Dequeue(job interfaces.Job) (*interfaces.Job, error) {
	// get the next job
	// if queue is empty return error
	// otherwise loop till the job has status not equal to cancelled
	for {
		select {
		case job := <-q.jobs:
			if q.jobsStatus[job.GetID()] == cancelled {
				continue
			}

			delete(q.jobsStatus, job.GetID())
			return &job, nil
		default:
			return nil, errors.New("job queue is empty")
		}
	}
}

func (q *JobQueue) CancelJob(jobID interfaces.JobID) error {
	// check if the job is currently in the queue
	// checking to avoid memory leak
	if _, ok := q.jobsStatus[jobID]; ok {
		q.jobsStatus[jobID] = cancelled
		return nil
	}

	return fmt.Errorf("job with id %v is not in the queue", jobID)
}
