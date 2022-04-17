package jobq

import "printer/persistence/model"

type JobIDGenerator struct {
	nextJobID model.JobID
}

func newJobIDGenerator() *JobIDGenerator {
	return &JobIDGenerator{
		nextJobID: 0,
	}
}

func (g *JobIDGenerator) newJobID() model.JobID {
	jobID := g.nextJobID
	g.nextJobID++

	return jobID
}
