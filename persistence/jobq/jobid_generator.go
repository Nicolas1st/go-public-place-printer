package jobq

type JobID uint

type JobIDGenerator struct {
	nextJobID JobID
}

func NewJobIDGenerator() *JobIDGenerator {
	return &JobIDGenerator{
		nextJobID: 0,
	}
}

func (g *JobIDGenerator) newJobID() JobID {
	jobID := g.nextJobID
	g.nextJobID++

	return jobID
}
