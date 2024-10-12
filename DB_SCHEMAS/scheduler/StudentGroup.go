package scheduler

// StudentGroup represents a group of students with their schedule details
type StudentGroup struct {
	ID        int
	Name      string
	Subjects  []string
	NoSubjects int
	TeacherIDs []int
	Hours     []int
}

// NewStudentGroup creates a new StudentGroup instance with default values
func NewStudentGroup() *StudentGroup {
	return &StudentGroup{
		Subjects:  make([]string, 10),
		Hours:     make([]int, 10),
		TeacherIDs: make([]int, 10),
	}
}

// SetID sets the ID of the StudentGroup
func (sg *StudentGroup) SetID(id int) {
	sg.ID = id
}

// GetID returns the ID of the StudentGroup
func (sg *StudentGroup) GetID() int {
	return sg.ID
}

// SetName sets the Name of the StudentGroup
func (sg *StudentGroup) SetName(name string) {
	sg.Name = name
}

// GetName returns the Name of the StudentGroup
func (sg *StudentGroup) GetName() string {
	return sg.Name
}

// SetSubjects sets the Subjects of the StudentGroup
func (sg *StudentGroup) SetSubjects(subjects []string) {
	sg.Subjects = subjects
}

// GetSubjects returns the Subjects of the StudentGroup
func (sg *StudentGroup) GetSubjects() []string {
	return sg.Subjects
}

// SetNoSubjects sets the number of subjects for the StudentGroup
func (sg *StudentGroup) SetNoSubjects(noSubjects int) {
	sg.NoSubjects = noSubjects
}

// GetNoSubjects returns the number of subjects for the StudentGroup
func (sg *StudentGroup) GetNoSubjects() int {
	return sg.NoSubjects
}

// SetTeacherIDs sets the TeacherIDs for the StudentGroup
func (sg *StudentGroup) SetTeacherIDs(teacherIDs []int) {
	sg.TeacherIDs = teacherIDs
}

// GetTeacherIDs returns the TeacherIDs for the StudentGroup
func (sg *StudentGroup) GetTeacherIDs() []int {
	return sg.TeacherIDs
}

// SetHours sets the Hours for the StudentGroup
func (sg *StudentGroup) SetHours(hours []int) {
	sg.Hours = hours
}

// GetHours returns the Hours for the StudentGroup
func (sg *StudentGroup) GetHours() []int {
	return sg.Hours
}
