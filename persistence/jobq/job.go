package jobq

type Job struct {
	ID       JobID
	FilePath string
	Username string
}

type JobBuilder func(id JobID, filepath, username string) *Job

func NewJobBuilder() JobBuilder {
	idGenerator := NewJobIDGenerator()

	return func(id JobID, filepath, username string) *Job {
		return &Job{
			ID:       idGenerator.newJobID(),
			FilePath: filepath,
			Username: username,
		}
	}
}
