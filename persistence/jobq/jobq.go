package jobq

type JobQueue struct {
	JobIDGenerator *JobIDGenerator
	queue          chan *Job
	jobsList       map[JobID]*Job // it's a map to make removal constant time
}

func NewJobQueue() *JobQueue {
	return &JobQueue{
		JobIDGenerator: newJobIDGenerator(),
		queue:          make(chan *Job, 20),
		jobsList:       map[JobID]*Job{},
	}
}

// Enqueue - adds job to the queue
func (q *JobQueue) Enqueue(job *Job) JobID {
	jobID := q.JobIDGenerator.newJobID()
	job.SetID(jobID)

	// push the job on to the queue
	q.queue <- job

	// store job for viewing
	q.jobsList[jobID] = job

	return jobID
}

// Dequeue - returns a non empty job, blocks execution when called, if not jobs are available
func (q *JobQueue) Dequeue() *Job {
	for {
		job := <-q.queue
		if job.Status != StatusCancelled {
			delete(q.jobsList, job.ID)
			return job
		}
	}
}

// Cancel - cancels job
func (q *JobQueue) CancelJob(jobID JobID) {
	job, ok := q.jobsList[jobID]
	if ok {
		job.CancelJob()
	}
}

func (q *JobQueue) GetAllJobs() []*Job {
	jobs := []*Job{}
	for _, v := range q.jobsList {
		jobs = append(jobs, v)
	}

	return jobs
}
