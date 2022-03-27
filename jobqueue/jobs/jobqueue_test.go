package jobqueue

import (
	"errors"
	"printer/job"
	"testing"
)

func TestEnqueueAndDequeque(t *testing.T) {

	// create job for testing
	job := job.NewJob(
		func() error {
			// do stuff
			return errors.New("error occured")
		},
	)

	// create queue
	q := NewJobQueue()

	// dequeueing from an empty queue
	_, err := q.Dequeue()
	if err == nil {
		t.Log("Dequeue from an empty queue must result in error")
		t.Fail()
	} else {
		t.Log("Dequeing from empty array was successful")
	}

	// enqueueing the job
	jobID := q.Enqueue(job)
	status, ok := q.jobsStatus[jobID]

	if !ok {
		t.Log("When enqueued the job's status must be stored in jobsStatus")
		t.Fail()
	}

	if status != toBeDone {
		t.Log("Just enqueued job must have the status equal to toBeDone")
		t.Fail()
	}

	// dequeueing the job
	{
		job, err := q.Dequeue()

		if err != nil {
			t.Log("Dequeueing from a non empty array should not result in an error")
			t.Fail()
		}

		if job.GetID() != jobID {
			t.Log("ID of enuqued job should not be changed when dequeued")
			t.Fail()
		}

	}

	// check the queue is empty
	{
		_, err := q.Dequeue()

		if err == nil {
			t.Log("Dequeue from an empty queue must result in error")
			t.Fail()
		}
	}

	// check if the job's status is removed
	if _, ok := q.jobsStatus[jobID]; ok {
		t.Log("The status of the job has not been removed, after dequing")
		t.Fail()
	}

}
