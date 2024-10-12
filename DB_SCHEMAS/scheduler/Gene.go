package scheduler

import (
	"math/rand"
	"time"
)

// Gene represents a permutation of slots as a timetable for a single student group
type Gene struct {
	SlotNo []int
	Days   int
	Hours  int
}

// NewGene creates a new Gene with a given index
func NewGene(i int) *Gene {
	days := inputData.DaysPerWeek
	hours := inputData.HoursPerDay
	rand.Seed(time.Now().UnixNano()) // Seed random number generator
	flag := make([]bool, days*hours)
	slotNo := make([]int, days*hours)

	for j := 0; j < days*hours; j++ {
		var rnd int
		for {
			rnd = rand.Intn(days * hours)
			if !flag[rnd] {
				break
			}
		}
		flag[rnd] = true
		slotNo[j] = i*days*hours + rnd
	}

	return &Gene{
		SlotNo: slotNo,
		Days:   days,
		Hours:  hours,
	}
}

// DeepClone creates a deep copy of the Gene instance
func (g *Gene) DeepClone() *Gene {
	// Create a new Gene instance and copy the data
	clone := *g
	clone.SlotNo = make([]int, len(g.SlotNo))
	copy(clone.SlotNo, g.SlotNo)
	return &clone
}
