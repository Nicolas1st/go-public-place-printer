package jobq

import "testing"

func TestJobIDUniqness(t *testing.T) {
	gen := NewJobIDGenerator()

	usedIDs := map[JobID]bool{}
	for i := 0; i < 10; i++ {
		id := gen.newJobID()

		if _, ok := usedIDs[id]; ok {
			t.Fatal("The id was used before")
		} else {
			usedIDs[id] = true
		}
	}
}
