package dbfiles

import (
	"log"
	"database/sql"
	_ "github.com/lib/pq"
	"strconv"
)
func CreateFacultyCourse(){
	// Create a new database connection
	db, err := sql.Open("postgres", "user=postgres dbname=mydb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	db.Exec(`
	CREATE TABLE IF NOT EXISTS FACULTY_COURSE
	 (
	   faculty_course_id VARCHAR(10), 
	   faculty_id  VARCHAR(10),
	   course_id VARCHAR(10), 
	   PRIMARY KEY (faculty_course_id),
	   CONSTRAINT fk_faculty_course 
	      FOREIGN KEY (faculty_id) REFERENCES FACULTY(faculty_id),
	   CONSTRAINT fk_course_faculty_course
	      FOREIGN KEY (course_id) REFERENCES COURSES(course_id)

	);`)
	
	// Close the database connection when we're done with it
	defer db.Close()
}
func InsertIntoFacultyCourse(faculty_id string , course_id string ){
	var faculty_course_id string = ""
	var faculty_course_count int = lengthOfDatabase("FACULTY_COURSE")
	 if faculty_course_count < 10 {
		 faculty_course_id = "FCREC0"+strconv.Itoa(faculty_course_count)
	 }else {
		 faculty_course_id = "FCREC"+strconv.Itoa(faculty_course_count)
	 }
	 faculty_course_count++
	db , err:= sql.Open("postgres","user=postgres dbname=mydb sslmode=disable")
	if err!=nil {
		log.Fatal(err)
	}
	defer db.Close()
	_,first_err := db.Exec("INSERT INTO FACULTY_COURSE (faculty_course_id , faculty_id , course_id   ) VALUES ($1,$2,$3)",faculty_course_id,faculty_id ,course_id)
	if first_err!=nil {
		log.Fatal(err)
	}
}
