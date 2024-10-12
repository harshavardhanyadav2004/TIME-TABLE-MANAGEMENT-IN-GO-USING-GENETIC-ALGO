package scheduler

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"strconv"
	"DB_SCHEMAS/insert_data_folder"
)

type Chromosome struct {
	CrossoverRate float64
	MutationRate  float64
	Hours         int
	Days          int
	NostGrp       int
	Fitness       float64
	Point         int
	Gene          []Gene
}

func NewChromosome() *Chromosome {
	c := &Chromosome{
		CrossoverRate: inputData.CrossoverRate,
		MutationRate:  inputData.MutationRate,
		Hours:         inputData.HoursPerDay,
		Days:          inputData.DaysPerWeek,
		NostGrp:       inputData.NoStudentGroup,
		Gene:          make([]Gene, inputData.NoStudentGroup),
	}

	for i := 0; i < inputData.NoStudentGroup; i++ {
		c.Gene[i] = *NewGene(i)
	}
	c.Fitness = c.GetFitness()
	return c
}

func (c *Chromosome) DeepClone() *Chromosome {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(c)
	if err != nil {
		return nil
	}
	
	dec := gob.NewDecoder(&buf)
	clone := &Chromosome{}
	err = dec.Decode(clone)
	if err != nil {
		return nil
	}
	return clone
}

func (c *Chromosome) GetFitness() float64 {
	c.Point = 0
	var teacherList = make(map[int]bool)
	TimeTableSlot := NewTimeTable().ReturnSlots()
	for i := 0; i < c.Hours*c.Days; i++ {
		teacherList = make(map[int]bool)
		for j := 0; j < c.NostGrp; j++ {
			slotIndex := c.Gene[j].SlotNo[i]
			slot := TimeTableSlot[slotIndex]
			if slot.StudentGroup != nil {
				if teacherList[slot.TeacherID] {
					c.Point++
				} else {
					teacherList[slot.TeacherID] = true
				}
			}
		}
	}
	c.Fitness = 1 - (float64(c.Point) / float64((c.NostGrp-1)*c.Hours*c.Days))
	c.Point = 0
	return c.Fitness
}
var courseIDToSubject = map[string]string{
    "AC-1101": "JS",
    "ME-1111": "MACHINE_LEARNING",
    "EVS-1101": "JIRA",
    "AM-1101": "DBMS",
    "PC-1101": "C",
    "AP-1101": "COMPUTER_ORGANISATION",
}
var total_time_table_map = make(map[string]map[string][]string)
var days = []string{"Mon","Tue","Wed","Thurs","Fri"}

func (c *Chromosome) PrintTimeTable() {
	TimeTableSlot := NewTimeTable().ReturnSlots()
	for i := 0; i < c.NostGrp; i++ {
		var batch_name string = ""
		status := false   
		l := 0
		for !status {
			slotIndex := c.Gene[i].SlotNo[l]
			slot := TimeTableSlot[slotIndex]
			if slot.StudentGroup != nil {
				fmt.Printf("Batch %s Timetable:\n", slot.StudentGroup.Name)
				batch_name = slot.StudentGroup.Name
				total_time_table_map[slot.StudentGroup.Name]  = make(map[string][]string)
				status = true
			}
			l++
		}
		fmt.Println(c.Days,c.Hours)
		for j := 0; j < c.Days; j++ {
			for k := 0; k < c.Hours; k++ {
				slotIndex := c.Gene[i].SlotNo[k+j*c.Hours]
				
				slot := TimeTableSlot[slotIndex]
				if slot.StudentGroup != nil {
					fmt.Printf("%s - %d  ", courseIDToSubject[slot.Subject],slot.TeacherID)
					var value = courseIDToSubject[slot.Subject]+"-"+strconv.Itoa(slot.TeacherID)
					total_time_table_map[batch_name][days[j]] = append(total_time_table_map[batch_name][days[j]],value )
					//fmt.Println(total_time_table_map)
				} else {
					fmt.Print("Lab Period ")
					total_time_table_map[batch_name][days[j]] = append(total_time_table_map[batch_name][days[j]],"Lab Period")
					//fmt.Println(total_time_table_map)
				}
			
			}
			fmt.Println()
		}
		fmt.Println()
	}
	fmt.Println(total_time_table_map)
	insertdatafolder.InsertIntoDirectlyIntoHost(total_time_table_map)
	fmt.Println("Successfully Table was pushed")

}

func (c *Chromosome) PrintChromosome() {
	for i := 0; i < c.NostGrp; i++ {
		for j := 0; j < c.Hours*c.Days; j++ {
			fmt.Printf("%d ", c.Gene[i].SlotNo[j])
		}
		fmt.Println()
	}
}

func (c *Chromosome) CompareTo(other *Chromosome) int {
	if c.Fitness == other.Fitness {
		return 0
	} else if c.Fitness > other.Fitness {
		return -1
	} else {
		return 1
	}
}

