package jobq

import (
	"printer/persistence/model"
)

type JobQueue struct {
	JobIDGenerator *JobIDGenerator
	queue          chan *model.Job
	jobsList       map[model.JobID]*model.Job // it's a map to make removal constant time
}

func NewJobQueue() *JobQueue {
	return &JobQueue{
		JobIDGenerator: newJobIDGenerator(),
		queue:          make(chan *model.Job, 20),
		jobsList:       map[model.JobID]*model.Job{},
	}
}

// Enqueue - adds job to the queue
func (q *JobQueue) Enqueue(job *model.Job) model.JobID {
	jobID := q.JobIDGenerator.newJobID()
	job.SetID(jobID)

	// push the job on to the queue
	q.queue <- job

	// store job for viewing
	q.jobsList[jobID] = job

	return jobID
}

// Dequeue - returns a non empty job, blocks execution when called, if not jobs are available
func (q *JobQueue) Dequeue() *model.Job {
	for {
		job := <-q.queue
		if job.Status != model.StatusCancelled {
			delete(q.jobsList, job.ID)
			return job
		}
	}
}

// Cancel - cancels job
func (q *JobQueue) CancelJob(jobID model.JobID) {
	q.jobsList[jobID].CancelJob()
}

func (q *JobQueue) GetAllJobs() []*model.Job {
	jobs := []*model.Job{}
	for _, v := range q.jobsList {
		jobs = append(jobs, v)
	}

	return jobs
}
