package main 

import (
	"DB_SCHEMAS/dbfiles"
	"DB_SCHEMAS/scheduler"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)
func ReadFromTheFiles(file_name string) [][] string{
	//insert into all tables
	//Reading from the csv file 
	file,err:= os.Open(file_name)
	if err!=nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records,err:= reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Read successfully")
	return records
}
func FillDetails (){
	 fmt.Println("Creating the databases using the postgres")
	 //Creating the databases
	 dbfiles.CreateBranches()
	 dbfiles.CreateCourses()
	 dbfiles.CreateFaculty()
	 dbfiles.CreateStudents()
	 dbfiles.CreateFacultyCourse()
	 dbfiles.CreateClass()
	 dbfiles.CreateAttendance()
	 dbfiles.CreateEnrollments()
	 //Inserting into the branches table 

	 var records [][] string 
	 records=ReadFromTheFiles("C:\\Users\\Admin\\Desktop\\DB_SCHEMAS\\data_folders\\college_branches.csv")
	 fmt.Println(records)
	 for _,record := range records[1:] {
		dbfiles.InsertIntoTable(record[0])
	 }
	 fmt.Println("Inserted into the branches table successfully")

	 //Inserting into the Faculty table 
	records=ReadFromTheFiles("C:\\Users\\Admin\\Desktop\\DB_SCHEMAS\\data_folders\\faculty_details.csv")
	 for _,record := range records[1:]{
		dbfiles.InsertIntoFaculty(record[0],record[1],record[2])

	 }
	 fmt.Println("Inserted into the faculty")

	 //Inserting into the course_details 
	 records=ReadFromTheFiles("C:\\Users\\Admin\\Desktop\\DB_SCHEMAS\\data_folders\\course_details.csv")
	 for _,record := range records[1:]{
		dbfiles.InsertIntoCourses(record[0],record[1],7)

	 }
	 fmt.Println("Inserted into the courses table successfully") 

	 //Inserting into the faculty_course
	 records=ReadFromTheFiles("C:\\Users\\Admin\\Desktop\\DB_SCHEMAS\\data_folders\\faculty_course_details.csv")
	 for _,record := range records[1:]{
		dbfiles.InsertIntoFacultyCourse(record[0],record[1])

	 }
	 fmt.Println("Inserted into the faculty_courses table successfully")
}
func main(){

	scheduler.CallFunction()
}