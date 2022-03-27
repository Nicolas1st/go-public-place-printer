package jobqueue

import (
	"printer/interfaces"
	"testing"
)

func TestNewJobIDGenerator(t *testing.T) {

	var testsReturns = []interfaces.JobID{1, 2, 3, 4, 5}

	newUniqueID := NewJobIDGenerator()

	for _, expected := range testsReturns {
		id := newUniqueID()
		t.Logf("Expected: %v, Recieved: %v", expected, id)
		if id != expected {
			t.Fatal("Failed")
		}
	}

}
