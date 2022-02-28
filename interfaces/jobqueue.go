package interfaces

type JobQueue interface {
	Enqueue(Job) JobID
	Dequeue() (*Job, error)
	CancelJob(JobID) error
}
