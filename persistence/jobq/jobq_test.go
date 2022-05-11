package jobq

import (
	"printer/persistence/model"
	"testing"
)

func TestEnqueueAndDequeque(t *testing.T) {
	// create job for testing
	jobIDGenerator := newJobIDGenerator()
	id := jobIDGenerator.newJobID()

	job := model.NewJob("/path/stuff", "stuff", "username")
	job.SetID(id)

	q := NewJobQueue()

	// enqueueing the job
	jobID := q.Enqueue(job)
	if job.Status != model.StatusToBeDone {
		t.Log("The job must have the status of to be done")
		t.Fail()
	}

	// dequeueing the job
	{
		job := q.Dequeue()

		if job.ID != jobID {
			t.Log("ID of enuqued job should not be changed when dequeued")
			t.Fail()
		}
	}

	// check if the job's status has been removed from the list
	_, ok := q.jobsList[jobID]
	if ok {
		t.Log("Job must be removed from the list, after it was dequeued")
		t.Fail()
	}

}
