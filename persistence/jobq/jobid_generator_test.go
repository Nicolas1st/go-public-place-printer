package jobq

import (
	"printer/persistence/model"
	"testing"
)

func TestJobIDUniqness(t *testing.T) {
	gen := newJobIDGenerator()

	usedIDs := map[model.JobID]bool{}
	for i := 0; i < 10; i++ {
		id := gen.newJobID()

		if _, ok := usedIDs[id]; ok {
			t.Fatal("The id was used before")
		} else {
			usedIDs[id] = true
		}
	}
}
