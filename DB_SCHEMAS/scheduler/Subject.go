package scheduler

// Subject represents a subject with its ID, name, and associated teachers
type Subject struct {
	ID         int
	Name       string
	Teachers   []Teacher
	NoTeachers int
}

// NewSubject creates a new Subject instance with default values
func NewSubject() *Subject {
	return &Subject{
		Teachers: make([]Teacher, 20), // Initial capacity for 20 teachers
	}
}
