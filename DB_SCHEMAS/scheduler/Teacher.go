package scheduler

// Teacher represents a teacher with their ID, name, subject, and number of assignments
type Teacher struct {
	ID       int
	Name     string
	Subject  string
	Assigned int
}

// NewTeacher creates a new Teacher instance with default values
func NewTeacher() *Teacher {
	return &Teacher{}
}
