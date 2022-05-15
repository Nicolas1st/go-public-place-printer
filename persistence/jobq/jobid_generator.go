package jobq

type JobIDGenerator struct {
	nextJobID JobID
}

func newJobIDGenerator() *JobIDGenerator {
	return &JobIDGenerator{
		nextJobID: 0,
	}
}

func (g *JobIDGenerator) newJobID() JobID {
	jobID := g.nextJobID
	g.nextJobID++

	return jobID
}
