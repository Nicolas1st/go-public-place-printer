package jobq

import "printer/interfaces"

func NewJobIDGenerator() func() interfaces.JobID {
	var startID interfaces.JobID = 0
	return func() interfaces.JobID {
		startID += 1
		return startID
	}
}
