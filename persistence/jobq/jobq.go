package jobq

import (
	"errors"
	"fmt"
)

type jobStatus byte

const (
	cancelled jobStatus = iota
	toBeDone
)

type JobQueue struct {
	JobIDGenerator *JobIDGenerator
	jobs           chan Job
	jobsStatus     map[JobID]jobStatus // to avoid linear search time
}

func NewJobQueue() *JobQueue {
	return &JobQueue{
		JobIDGenerator: NewJobIDGenerator(),
		jobs:           make(chan Job, 20),
		jobsStatus:     make(map[JobID]jobStatus),
	}
}

func (q *JobQueue) Enqueue(job Job) JobID {
	jobID := job.ID

	// set the status for the job
	q.jobsStatus[jobID] = toBeDone

	// push the job on to the queue
	q.jobs <- job

	return jobID
}

func (q *JobQueue) Dequeue() (Job, error) {
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
			return Job{}, errors.New("job queue is empty")
		}
	}
}

func (q *JobQueue) CancelJob(jobID JobID) error {
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
