package jobq

type JobIDGenerator struct {
	nextJobID int
}

func NewJobIDGenerator() *JobIDGenerator {
	return &JobIDGenerator{
		nextJobID: 0,
	}
}

func (g *JobIDGenerator) newJobID() int {
	jobID := g.nextJobID
	g.nextJobID++

	return jobID
}
