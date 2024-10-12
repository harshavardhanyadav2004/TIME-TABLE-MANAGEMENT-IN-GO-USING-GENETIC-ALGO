package dbfiles

import (
	"database/sql"
	"log"
	"strconv"

	_ "github.com/lib/pq"
)
func CreateAttendance(){
	db,err:=sql.Open("postgres","user=postgres dbname=mydb sslmode=disable")
	if err != nil {
		 log.Fatal(err)
	}
	db.Exec(`
	CREATE TABLE IF NOT EXISTS attendance (
	attendance_id VARCHAR(100),
	student_id VARCHAR(10),
	class_id VARCHAR(10),
	date DATE , 
	status VARCHAR(20),
	PRIMARY KEY (attendance_id),
	CONSTRAINT fk_attendance 
	   FOREIGN KEY (student_id) REFERENCES STUDENTS(student_id),
	CONSTRAINT fk_attendance_course
	   FOREIGN KEY (class_id) REFERENCES CLASS(class_id)
	);`)
	defer db.Close()
}
func InsertIntoAttendance(student_id string , class_id string , date string , status string ){
	var attendance_id string = ""
     var attendance_count int = lengthOfDatabase("attendance")
	if attendance_count < 10 {
		 attendance_id = "AAIREC"+strconv.Itoa(attendance_count)
	}else {
		 attendance_id = "AAIREC"+strconv.Itoa(attendance_count)
	}
	attendance_count++
	db,err:=sql.Open("postgres","user=postgres dbname=mydb sslmode=disable")
	if err !=nil {
		log.Fatal(err)
	}
	_ , new_err := db.Exec(
		`INSERT INTO attendance (attendance_id , student_id , class_id , date ,status)
		 VALUES (?,?,?,?,?)
		`,
		attendance_id ,
		student_id ,
		class_id ,
		date ,
		status)
	if new_err != nil {
		log.Fatal(new_err)
	}
	defer db.Close()
}
