package scheduler

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// InputData holds the data for student groups and teachers
// NewTimeTable creates a new TimeTable and initializes slots
type InputData struct {
	StudentGroup []StudentGroup
	Teacher      []Teacher
	CrossoverRate float64
	MutationRate  float64
	NoStudentGroup int
	NoTeacher      int
	HoursPerDay    int
	DaysPerWeek    int
}

// NewInputData initializes the InputData struct
func NewInputData() *InputData {
	return &InputData{
		CrossoverRate: 1.0,
		MutationRate:  0.1,
		HoursPerDay:   7,
		DaysPerWeek:   5,
	}
}

// ClassFormat checks if the line has the correct format
func (data *InputData) ClassFormat(l string) bool {
	st := strings.Fields(l)
	return len(st) == 3
}

// TakeInput reads input data from a file
func (data *InputData) TakeInput() {
	file, err := os.Open("C:\\Users\\Admin\\Desktop\\ClassRoom_Management\\DB_SCHEMAS\\scheduler\\input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "studentgroups" {
			data.StudentGroup = []StudentGroup{}
			for scanner.Scan() {
				line = scanner.Text()
				if line == "teachers" {
					break
				}
				st := strings.Fields(line)
				sg := StudentGroup{
					ID:         len(data.StudentGroup),
					Name:       st[0],
					Subjects:   make([]string, (len(st)-1)/2),
					Hours:      make([]int, (len(st)-1)/2),
					TeacherIDs: make([]int, (len(st)-1)/2),
				}
				for k := 1; k < len(st); k += 2 {
					sg.Subjects[(k-1)/2] = st[k]
					hours, _ := strconv.Atoi(st[k+1])
					sg.Hours[(k-1)/2] = hours
					sg.NoSubjects++
				}
				data.StudentGroup = append(data.StudentGroup, sg)
			}
			data.NoStudentGroup = len(data.StudentGroup)
		}

		if line == "teachers" {
			data.Teacher = []Teacher{}
			for scanner.Scan() {
				line = scanner.Text()
				if line == "end" {
					break
				}
				st := strings.Fields(line)
				t := Teacher{
					ID:       len(data.Teacher),
					Name:     st[0],
					Subject:  st[1],
					Assigned: 0,
				}
				data.Teacher = append(data.Teacher, t)
			}
			data.NoTeacher = len(data.Teacher)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	data.AssignTeacher()
}

// AssignTeacher assigns a teacher to each subject for every student group
func (data *InputData) AssignTeacher() {
	for i := 0; i < data.NoStudentGroup; i++ {
		for j := 0; j < data.StudentGroup[i].NoSubjects; j++ {
			teacherID := -1
			assignedMin := -1

			subject := data.StudentGroup[i].Subjects[j]

			for k := 0; k < data.NoTeacher; k++ {
				if data.Teacher[k].Subject == subject {
					if assignedMin == -1 {
						assignedMin = data.Teacher[k].Assigned
						teacherID = k
					} else if assignedMin > data.Teacher[k].Assigned {
						assignedMin = data.Teacher[k].Assigned
						teacherID = k
					}
				}
			}

			data.Teacher[teacherID].Assigned++
			data.StudentGroup[i].TeacherIDs[j] = teacherID
		}
	}
}
var inputData = NewInputData()
func CallFunction(){
	// Call the function here
	inputData.TakeInput()
	NewSchedulerMain()
	finalSon.PrintTimeTable()

}
