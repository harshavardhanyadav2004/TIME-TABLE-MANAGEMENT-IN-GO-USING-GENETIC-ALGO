package dbfiles

import (
	"database/sql"
	"log"
	"strconv"

	_ "github.com/lib/pq"
)
func CreateEnrollments(){
	db, err := sql.Open("postgres", "user=postgres dbname=mydb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	db.Exec(`
	CREATE TABLE ENROLLMENTS (
	enrollment_id VARCHAR(10),
	student_id VARCHAR(10),
	class_id VARCHAR(10),
	enrollment_date DATE ,
	PRIMARY KEY (enrollment_id),
	CONSTRAINT FK_ENROLLMENTS 
	   FOREIGN KEY (student_id) REFERENCES STUDENTS(student_id),
	CONSTRAINT FK_ENROLLMENTS_2
	   FOREIGN KEY (class_id) REFERENCES CLASS(class_id)
	)
	`)
	defer db.Close()
}
func InsertIntoEnrollments(student_id string , class_id string , enrollment_date string){
	var enrollment_id string = ""
	var enrollment_count int = lengthOfDatabase("ENROLLMENTS")
	if enrollment_count < 10 {
		enrollment_id = "EIREC0"+strconv.Itoa(enrollment_count)
	}else {
		enrollment_id = "EICREC"+strconv.Itoa(enrollment_count)
	}
	enrollment_count++
	  db,err:=sql.Open("postgres","user=postgres dbname=mydb sslmode=disable")
	  if err != nil {
		log.Fatal(err)
	  }
	  _,new_err := db.Exec(
		`
		INSERT INTO ENROLLMENTS 
		(enrollment_id,student_id ,class_id, enrollment_date)
		VALUES (?,?,?,?)
		`,
		enrollment_id , student_id , class_id ,enrollment_date)
		if new_err !=nil {
			log.Fatal(err)
		}
		defer db.Close()
}