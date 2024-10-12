package scheduler

// TimeTable holds and generates slots for the timetable
type TimeTable struct {
	Slots []Slot
}

func NewTimeTable() *TimeTable {
	hours:=inputData.HoursPerDay
	days:=inputData.DaysPerWeek
	nostgrp:=inputData.NoStudentGroup
	slot := make([]Slot, hours*days*nostgrp)
	k := 0
	subjectno := 0
	hourcount := 1

	// Create as many slots as the number of blocks in the overall timetable
	for i := 0; i < nostgrp; i++ {
		subjectno = 0
		// For every slot in a week for a student group
		for j := 0; j < hours*days; j++ {
			sg := inputData.StudentGroup[i]

			// If all subjects have been assigned required hours, give free periods
			if subjectno >= sg.NoSubjects {
				slot[k] = Slot{} // Assuming zero value represents free period
				k++
			} else {
				slot[k] = Slot{
					StudentGroup: &sg,
					TeacherID:    sg.TeacherIDs[subjectno],
					Subject:      sg.Subjects[subjectno],
				}
				k++

				// If the required hours for the subject are not met, keep adding
				if hourcount < sg.Hours[subjectno] {
					hourcount++
				} else {
					hourcount = 1
					subjectno++
				}
			}
		}
	}

	return &TimeTable{
		Slots: slot,
	}
}

// ReturnSlots returns the slots
func (tt *TimeTable) ReturnSlots() []Slot {
	return tt.Slots
}
