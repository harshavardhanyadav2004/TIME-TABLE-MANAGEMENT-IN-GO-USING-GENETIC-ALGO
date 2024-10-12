package scheduler

// Slot represents a single block of the timetable
type Slot struct {
	StudentGroup *StudentGroup
	TeacherID    int
	Subject      string
}

// NewSlot creates a new Slot instance with the given parameters
func NewSlot(studentGroup *StudentGroup, teacherID int, subject string) *Slot {
	return &Slot{
		StudentGroup: studentGroup,
		TeacherID:    teacherID,
		Subject:      subject,
	}
}

// NewFreeSlot creates a Slot instance for a free period
func NewFreeSlot() *Slot {
	return &Slot{}
}
