package interfaces

type JobID uint

type Job interface {
	GetID() JobID
	SetID(JobID)
	Execute() error
}
