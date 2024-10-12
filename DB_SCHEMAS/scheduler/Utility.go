package scheduler

import "fmt"

// Utility provides static methods for printing input data and slots
type Utility struct{}

// PrintInputData prints the input data
func (Utility) PrintInputData() {
	fmt.Printf("Nostgrp=%d Noteachers=%d daysperweek=%d hoursperday=%d\n",
		inputData.NoStudentGroup,
		inputData.NoTeacher,
		inputData.DaysPerWeek,
		inputData.HoursPerDay,
	)

	for _, sg := range inputData.StudentGroup {
		fmt.Printf("%d %s\n", sg.ID, sg.Name)
		for j := 0; j < sg.NoSubjects; j++ {
			fmt.Printf("%s %d hrs %d\n", sg.Subjects[j], sg.Hours[j], sg.TeacherIDs[j])
		}
		fmt.Println("")
	}

	for _, t := range inputData.Teacher {
		fmt.Printf("%d %s %s %d\n", t.ID, t.Name, t.Subject, t.Assigned)
	}
}

// PrintSlots prints the slots
func (Utility) PrintSlots(tt *TimeTable) {
	days := inputData.DaysPerWeek
	hours := inputData.HoursPerDay
	nostgrp := inputData.NoStudentGroup

	fmt.Println("----Slots----")
	for i := 0; i < days*hours*nostgrp; i++ {
		if i < len(tt.Slots) && (tt.Slots[i] != Slot{}) {
			slot := tt.Slots[i]
			fmt.Printf("%d- %s %s %d\n", i, slot.StudentGroup.Name, slot.Subject, slot.TeacherID)
		} else {
			fmt.Println("Free Period")
		}
		if (i+1)%(hours*days) == 0 {
			fmt.Println("******************************")
		}
	}
}
